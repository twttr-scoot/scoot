package main

// Command line tool to enable testing for Bazel requests in Scoot
// We use this for client-side operations Scoot does not implement,
// and are difficult to reproduce using generic tools like grpc_cli.
// Not part of production deployment
//
// Supports subcommands (--help for usage):
// * remoteexecution.Command protobuf data structure insertion to a CAS
// * longrunning.GetOperation polling of operation/scootjob by name and pretty print of result

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	remoteexecution "github.com/twitter/scoot/bazel/remoteexecution"

	"github.com/twitter/scoot/bazel"
	"github.com/twitter/scoot/bazel/cas"
	"github.com/twitter/scoot/bazel/execution"
	"github.com/twitter/scoot/common/dialer"
	"github.com/twitter/scoot/common/log/hooks"
	scootproto "github.com/twitter/scoot/common/proto"
	"github.com/twitter/scoot/scootapi"
)

var uploadCmdStr string = "upload_command"
var uploadActionStr string = "upload_action"
var execCmdStr string = "execute"
var getOpCmdStr string = "get_operation"
var supportedCommands map[string]bool = map[string]bool{
	uploadCmdStr:    true,
	uploadActionStr: true,
	execCmdStr:      true,
	getOpCmdStr:     true,
}

func main() {
	log.AddHook(hooks.NewContextHook())

	// TODO additional commands
	// - upload Directory (tbd)
	// - download outputs, stdout/err

	// Subcommand flag definitions

	// Upload Command
	uploadCommand := flag.NewFlagSet(uploadCmdStr, flag.ExitOnError)
	uploadAddr := uploadCommand.String("cas_addr", scootapi.DefaultApiBundlestore_GRPC, "'host:port' of grpc CAS server")
	uploadEnv := uploadCommand.String("env", "", "comma-separated command environment variables, i.e. \"key1=val1,key2=val2\"")
	uploadOutputFiles := uploadCommand.String("output_files", "", "Output files to ingest as comma-separated list: '/file1,/dir/file2'")
	uploadOutputDirs := uploadCommand.String("output_dirs", "", "Output dirs to ingest as comma-separated list: '/dir'")

	// Upload Action
	uploadAction := flag.NewFlagSet(uploadActionStr, flag.ExitOnError)
	actionAddr := uploadAction.String("cas_addr", scootapi.DefaultApiBundlestore_GRPC, "'host:port' of grpc CAS server")
	actionCommandDigest := uploadAction.String("command", "", "Command digest as '<hash>/<size>'")
	actionRootDigest := uploadAction.String("input_root", "", "Input root digest as '<hash>/<size>'")
	actionNoCache := uploadAction.Bool("no_cache", false, "Flag to prevent result caching")

	// Execute
	execCommand := flag.NewFlagSet(execCmdStr, flag.ExitOnError)
	execAddr := execCommand.String("grpc_addr", scootapi.DefaultSched_GRPC, "'host:port' of grpc Exec server")
	execActionDigest := execCommand.String("action", "", "Action digest as '<hash>/<size>'")
	execSkipCache := execCommand.Bool("skip_cache", false, "Skip checking for cached results")
	execJson := execCommand.Bool("json", false, "Print operation as JSON")

	// Get Operation
	getCommand := flag.NewFlagSet(getOpCmdStr, flag.ExitOnError)
	getAddr := getCommand.String("grpc_addr", scootapi.DefaultSched_GRPC, "'host:port' of grpc Exec server")
	getName := getCommand.String("name", "", "Operation name to query")
	getJson := execCommand.Bool("json", false, "Print operation as JSON")

	// Parse input flags
	if len(os.Args) < 2 {
		printSupported()
		os.Exit(1)
	}
	switch os.Args[1] {
	case uploadCmdStr:
		uploadCommand.Parse(os.Args[2:])
	case uploadActionStr:
		uploadAction.Parse(os.Args[2:])
	case execCmdStr:
		execCommand.Parse(os.Args[2:])
	case getOpCmdStr:
		getCommand.Parse(os.Args[2:])
	default:
		printSupported()
		os.Exit(1)
	}

	// Run parsed subcommand
	if uploadCommand.Parsed() {
		uploadArgv := uploadCommand.Args()
		if len(uploadArgv) == 0 {
			log.Fatalf("Argv required for %s - will interpret all non-flag arguments as Argv", uploadCmdStr)
		}
		uploadBzCommand(uploadArgv, *uploadAddr, *uploadEnv, *uploadOutputFiles, *uploadOutputDirs)
	} else if uploadAction.Parsed() {
		if *actionCommandDigest == "" || *actionRootDigest == "" {
			log.Fatalf("command and input_root required for %s", execCmdStr)
		}
		uploadBzAction(*actionAddr, *actionCommandDigest, *actionRootDigest, *actionNoCache)
	} else if execCommand.Parsed() {
		if *execActionDigest == "" {
			log.Fatalf("action digest required for %s", execCmdStr)
		}
		execute(*execAddr, *execActionDigest, *execSkipCache, *execJson)
	} else if getCommand.Parsed() {
		if *getName == "" {
			log.Fatalf("name required for %s", getOpCmdStr)
		}
		getOperation(*getAddr, *getName, *getJson)
	} else {
		log.Fatal("No expected commands parsed")
	}
}

func uploadBzCommand(cmdArgs []string, casAddr, env, outputFilesStr, outputDirsStr string) {
	envMap := make(map[string]string)
	for _, pair := range strings.Split(env, ",") {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			continue
		}
		envMap[kv[0]] = kv[1]
	}
	log.Infof("Using argv: %q env: %s", cmdArgs, envMap)

	// create Command struct from inputs
	cmdEnvVars := []*remoteexecution.Command_EnvironmentVariable{}
	for k, v := range envMap {
		cmdEnvVars = append(cmdEnvVars, &remoteexecution.Command_EnvironmentVariable{Name: k, Value: v})
	}

	outputFiles := []string{}
	outputDirs := []string{}
	for _, f := range strings.Split(outputFilesStr, ",") {
		if f == "" {
			continue
		}
		outputFiles = append(outputFiles, f)
	}
	for _, d := range strings.Split(outputDirsStr, ",") {
		if d == "" {
			continue
		}
		outputDirs = append(outputDirs, d)
	}

	cmd := &remoteexecution.Command{
		Arguments:            cmdArgs,
		EnvironmentVariables: cmdEnvVars,
		OutputFiles:          outputFiles,
		OutputDirectories:    outputDirs,
	}

	// serialize and get hash/size
	bytes, err := proto.Marshal(cmd)
	if err != nil {
		log.Fatalf("Error serializing command message: %s", err)
	}
	hash, size, err := scootproto.GetSha256(cmd)
	if err != nil {
		log.Fatalf("Error serializing command message: %s", err)
	}

	// upload command to CAS
	r := dialer.NewConstantResolver(casAddr)
	digest := &remoteexecution.Digest{Hash: hash, SizeBytes: size}
	err = cas.ByteStreamWrite(r, digest, bytes)
	if err != nil {
		log.Fatalf("Error writing to CAS: %s", err)
	}

	log.Info("Wrote to CAS successfully")
	fmt.Printf("%s/%d\n", hash, size)
}

func uploadBzAction(casAddr, commandDigestStr, rootDigestStr string, noCache bool) {
	commandDigest, err := bazel.DigestFromString(commandDigestStr)
	if err != nil {
		log.Fatalf("Error converting action to Digest: %s", err)
	}
	rootDigest, err := bazel.DigestFromString(rootDigestStr)
	if err != nil {
		log.Fatalf("Error converting action to Digest: %s", err)
	}

	action := &remoteexecution.Action{
		CommandDigest:   commandDigest,
		InputRootDigest: rootDigest,
		DoNotCache:      noCache,
	}

	// serialize and get hash/size
	bytes, err := proto.Marshal(action)
	if err != nil {
		log.Fatalf("Error serializing action message: %s", err)
	}
	hash, size, err := scootproto.GetSha256(action)
	if err != nil {
		log.Fatalf("Error serializing action message: %s", err)
	}

	// upload action to CAS
	r := dialer.NewConstantResolver(casAddr)
	digest := &remoteexecution.Digest{Hash: hash, SizeBytes: size}
	err = cas.ByteStreamWrite(r, digest, bytes)
	if err != nil {
		log.Fatalf("Error writing to CAS: %s", err)
	}

	log.Info("Wrote to CAS successfully")
	fmt.Printf("%s/%d\n", hash, size)
}

func execute(execAddr, actionDigestStr string, skipCache bool, execJson bool) {
	r := dialer.NewConstantResolver(execAddr)

	actionDigest, err := bazel.DigestFromString(actionDigestStr)
	if err != nil {
		log.Fatalf("Error converting action to Digest: %s", err)
	}

	operation, err := execution.Execute(r, actionDigest, skipCache)
	if err != nil {
		log.Fatalf("Error making Execute request: %s", err)
	}
	log.Info(execution.ExecuteOperationToStr(operation))
	if execJson {
		b, err := json.Marshal(operation)
		if err != nil {
			log.Fatalf("Error converting operation to JSON: %v", err)
		}
		fmt.Printf("%s\n", b)
	}
}

func getOperation(execAddr, opName string, getJson bool) {
	r := dialer.NewConstantResolver(execAddr)
	operation, err := execution.GetOperation(r, opName)
	if err != nil {
		log.Fatalf("Error making GetOperation request: %s", err)
	}

	log.Info(execution.ExecuteOperationToStr(operation))
	if getJson {
		b, err := json.Marshal(operation)
		if err != nil {
			log.Fatalf("Error converting operation to JSON: %v", err)
		}
		fmt.Printf("%s\n", b)
	}
}

func printSupported() {
	cmds := make([]string, 0, len(supportedCommands))
	for k := range supportedCommands {
		cmds = append(cmds, k)
	}
	fmt.Printf("Supported commands: %s\n", strings.Join(cmds, ", "))
}
