package scootapi

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Where is Cloud Scoot running?
// We store the answer (as host:port) in ~/.cloudscootaddr
// There are two lines: first is the sched thrift addr, second is the bundlestore http addr.
// TODO: this will eventually store only the thrift addr and http addr
//       for a single instance of apiserver, though several may be running.
// TODO: can we get rid of this and exclusively rely on a Fetcher to find instances?

// Get the path of the file containing the address for scootapi to use
func GetScootapiAddrPath() string {
	return path.Join(os.Getenv("HOME"), ".cloudscootaddr")
}

// Get the scootapi address (as host:port)
func GetScootapiAddr() (sched string, api string, err error) {
	data, err := ioutil.ReadFile(GetScootapiAddrPath())
	if err != nil {
		return "", "", err
	}
	addrs := strings.Split(string(data), "\n")
	if len(addrs) != 2 {
		return "", "", errors.New("Expected both sched and api addrs, got: " + string(data))
	}
	return string(addrs[0]), string(addrs[1]), nil
}

// Set the scootapi address (as host:port)
func SetScootapiAddr(sched string, api string) {
	ioutil.WriteFile(GetScootapiAddrPath(), []byte(sched+"\n"+api), 0777)
}

func APIAddrToBundlestoreURI(addr string) string {
	return "http://" + addr + "/bundle/"
}
