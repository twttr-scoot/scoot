package local

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/scootdev/scoot/os/temp"
	"github.com/scootdev/scoot/runner"
	"github.com/scootdev/scoot/runner/execer/execers"
	"github.com/scootdev/scoot/snapshot/snapshots"
)

type testEnv struct {
	wg   *sync.WaitGroup
	qr   *QueueingRunner
	size int
}

const errorMsgFromRunner = "Error in fakeRunner Run()"

// Send a run request that pauses, then another run request.
// Verify that the status of the first run request is running, and
// the status of the second request is queued.
// Then send a signal to allow the first (paused) run request to finish.
// Wait 1ms then verify that the second run request was run.
func TestQueueing2Messages(t *testing.T) {
	testEnv := setup(false, 10, t)
	defer teardown(testEnv)
	qr := testEnv.qr
	testEnv.wg.Add(1) // set the pause condition

	// send the first command - it should pause
	run1 := assertRun(t, qr, running(), "pause", "complete 0")
	run2 := assertRun(t, qr, pending(), "complete 1")
	testEnv.wg.Done() // send signal to end first (paused) request
	assertWait(t, qr, run1, complete(0), "n/a")
	assertWait(t, qr, run2, complete(1), "n/a")
}

func TestQueueingMoreThanMaxMessage(t *testing.T) {
	testEnv := setup(false, 4, t)
	defer teardown(testEnv)
	qr := testEnv.qr
	testEnv.wg.Add(1) // set the pause condition

	var runIDs []runner.RunId

	// block the runner
	runID := assertRun(t, qr, running(), "pause", "complete 0")
	runIDs = append(runIDs, runID)

	// fill the queue
	for i := 0; i < 4; i++ {
		runID := assertRun(t, qr, pending(), "complete 0")
		runIDs = append(runIDs, runID)
	}

	// overflow
	_, err := qr.Run(&runner.Command{Argv: []string{"complete 5"}})
	if err == nil || strings.Compare(QueueFullMsg, err.Error()) != 0 {
		t.Fatal("Should not be able to schedule: ", err)
	}

	// resume
	testEnv.wg.Done()

	// drain
	for _, id := range runIDs {
		assertWait(t, qr, id, complete(0), "n/a")
	}
	runIDs = nil

	// repeat
	testEnv.wg.Add(1)
	runID = assertRun(t, qr, running(), "pause", "complete 0")
	runIDs = append(runIDs, runID)
	for i := 0; i < 4; i++ {
		runID := assertRun(t, qr, pending(), "complete 0")
		runIDs = append(runIDs, runID)
	}
	_, err = qr.Run(&runner.Command{Argv: []string{"complete 5"}})
	if err == nil || strings.Compare(QueueFullMsg, err.Error()) != 0 {
		t.Fatal("Should not be able to schedule: ", err)
	}
	testEnv.wg.Done()
	for _, id := range runIDs {
		assertWait(t, qr, id, complete(0), "n/a")
	}
}

func TestUnknownRunIdInStatusRequest(t *testing.T) {
	testEnv := setup(false, 4, t)
	defer teardown(testEnv)
	qr := testEnv.qr

	assertWait(t, qr, runner.RunId("not a real run id"), badRequest(UnknownRunIdMsg), "n/a")
}

// func TestRunnerReturningAnErrorOnRunRequest(t *testing.T) {

// 	testEnv := setup(true, 4, t)
// 	defer teardown(testEnv)
// 	qr := testEnv.qr

// 	validateRunRequest("Run():", []string{"Command:", "Run", "return", "error", "test"}, []runner.ProcessState{runner.BADREQUEST}, errorMsgFromRunner, qr, t)
// }

// func SkipTestStatusAll(t *testing.T) {

// 	testEnv := setup(false, 10, t)
// 	defer teardown(testEnv)
// 	qr := testEnv.qr

// 	testEnv.wg.Add(1) // set the pause condition

// 	// put commands in the queue with the first command paused
// 	var args [6][]string

// 	args[0] = []string{"pause", "complete 0"} // send the first command - it should pause
// 	validateRunRequest("phase1, 1st Run():", args[0], []runner.ProcessState{runner.PENDING}, "", qr, t)
// 	waitForStatus("phase1, wait for first run running:", runner.RunId("0"), []runner.ProcessState{runner.RUNNING}, 10*time.Millisecond, qr, t)

// 	for i := 1; i < 5; i++ {
// 		// send a second command
// 		args[i] = []string{"complete " + strconv.Itoa(i)}
// 		tag := fmt.Sprintf("phase1, Run() %d:", i)
// 		validateRunRequest(tag, args[i], []runner.ProcessState{runner.PENDING}, "", qr, t)
// 	}

// 	// check the status
// 	statuss := qr.getStatusAll()

// 	if len(statuss) != 5 {
// 		t.Fatalf("Wrong length on statusAll response: expected 5, got %d", len(statuss))
// 	}

// 	if statuss[0].State != runner.RUNNING && statuss[0].State != runner.PREPARING && statuss[0].State != runner.PENDING {
// 		t.Fatalf("Wrong statuss[0]: expected Running, Preparing or Pending, got %s", statuss[0].State)
// 	}

// 	for i := 1; i < 5; i++ {
// 		if statuss[i].State != runner.PENDING {
// 			t.Fatalf("Wrong statuss[%d]: expected Pending, got %s", i, statuss[i].State)
// 		}
// 	}
// 	testEnv.wg.Done()

// }

// func SkipTestAbortTop2ReuqestsWhenPaused(t *testing.T) {
// 	testEnv := setup(false, 10, t)
// 	defer teardown(testEnv)
// 	qr := testEnv.qr

// 	testEnv.wg.Add(1) // set the pause condition

// 	// put commands in the queue with the first command paused
// 	var args [6][]string

// 	args[0] = []string{"pause", "complete 0"} // send the first command - it should pause
// 	validateRunRequest("phase1, 1st Run():", args[0], []runner.ProcessState{runner.PENDING}, "", qr, t)
// 	waitForStatus("phase1, wait for first run running:", runner.RunId("0"), []runner.ProcessState{runner.RUNNING}, 10*time.Millisecond, qr, t)

// 	for i := 1; i < 5; i++ {
// 		// send a second command
// 		args[i] = []string{"complete " + strconv.Itoa(i)}
// 		tag := fmt.Sprintf("phase1, Run() %d:", i)
// 		validateRunRequest(tag, args[i], []runner.ProcessState{runner.PENDING}, "", qr, t)
// 	}

// 	for i := 0; i < 2; i++ {
// 		s, err := qr.Abort(runner.RunId(strconv.Itoa(i)))

// 		if err != nil {
// 			t.Fatalf("Error aborting run(%d):'%s'", i, err.Error())
// 		}

// 		if s.State != runner.ABORTED {
// 			t.Fatalf("Aborted run(%d) did not return 'Aborted' state, got: '%s'", i, s.State)
// 		}

// 		s, err = qr.Status(runner.RunId(strconv.Itoa(i)))
// 		if err != nil {
// 			t.Fatalf("Error getting status of run(%d):'%s'", i, err.Error())
// 		}
// 		if s.State != runner.ABORTED {
// 			t.Fatalf("Error getting status of run(%d) expected state 'Aborted', got '%s'", i, s.State)
// 		}

// 	}

// 	testEnv.wg.Done()

// 	sAll, err := qr.StatusAll()

// 	if err != nil {
// 		t.Fatalf("Error getting status All '%s'", err.Error())
// 	}

// 	if sAll[0].State != runner.ABORTED {
// 		t.Fatalf("Error, statusAll state for run(0), expected 'Aborted', got '%s'", sAll[0].State)
// 	}
// 	if sAll[1].State != runner.ABORTED {
// 		t.Fatalf("Error, statusAll state for run(1), expected 'Aborted', got '%s'", sAll[1].State)
// 	}
// 	for i := 2; i < 5; i++ {
// 		if sAll[i].State == runner.ABORTED {
// 			t.Fatalf("Error, statusAll state for run(%d), expected anything but 'Aborted', got 'Aborted'", i)
// 		}
// 	}
// }

// func SkipTestAbortingFirst3RequestsInQueue(t *testing.T) {
// 	testEnv := setup(false, 10, t)
// 	defer teardown(testEnv)
// 	qr := testEnv.qr

// 	testEnv.wg.Add(1) // set the pause condition

// 	// put commands in the queue with the first command paused
// 	var args [6][]string

// 	args[0] = []string{"pause", "complete 0"} // send the first command - it should pause
// 	validateRunRequest("phase1, 1st Run():", args[0], []runner.ProcessState{runner.PREPARING}, "", qr, t)
// 	waitForStatus("phase1, wait for first run running:", runner.RunId("0"), []runner.ProcessState{runner.RUNNING}, 10*time.Millisecond, qr, t)

// 	for i := 1; i < 5; i++ {
// 		// send a second command
// 		args[i] = []string{"complete " + strconv.Itoa(i)}
// 		tag := fmt.Sprintf("phase1, Run() %d:", i)
// 		validateRunRequest(tag, args[i], []runner.ProcessState{runner.PENDING}, "", qr, t)
// 	}

// 	for i := 0; i < 3; i++ {
// 		s, err := qr.Abort(runner.RunId(strconv.Itoa(i)))

// 		if err != nil {
// 			t.Fatalf("Error aborting run (%d), %s", i, err.Error())
// 		}
// 		if s.State != runner.ABORTED {
// 			t.Fatalf("Error expected Aborted state for run(%d), got '%s'", i, s.State)
// 		}
// 	}

// 	// let the rest of the run requests run
// 	testEnv.wg.Done()

// 	sAll, err := qr.StatusAll()

// 	if err != nil {
// 		t.Fatalf("Error getting status All '%s'", err.Error())
// 	}

// 	for i := 3; i < 5; i++ {
// 		if sAll[i].State == runner.ABORTED {
// 			t.Fatalf("Error, statusAll state for run(%d), expected anything but 'Aborted', got 'Aborted'", i)
// 		}
// 	}

// }

func setup(useFakeRunner bool, size int, t *testing.T) *testEnv {

	runnerAvailableCh := make(chan struct{})

	ctx := context.TODO()

	var runner runner.Runner
	var wg *sync.WaitGroup
	if useFakeRunner {
		runner = nil
	} else {
		runner, wg = makeRunnerWithSimExecer(runnerAvailableCh, t)
	}

	qr := NewQueuingRunner(ctx, runner, size, runnerAvailableCh).(*QueueingRunner)

	return &testEnv{wg: wg, qr: qr, size: size}
}

func makeRunnerWithSimExecer(runnerAvailableCh chan struct{}, t *testing.T) (runner.Runner, *sync.WaitGroup) {
	wg := &sync.WaitGroup{}
	ex := execers.NewSimExecer(wg)
	tempDir, err := temp.TempDirDefault()
	if err != nil {
		t.Fatalf("Test setup() failed getting temp dir:%s", err.Error())
	}

	outputCreator, err := NewOutputCreator(tempDir)
	if err != nil {
		t.Fatalf("Test setup() failed getting output creator:%s", err.Error())
	}
	r := NewSimpleReportBackRunner(ex, snapshots.MakeInvalidCheckouter(), outputCreator, runnerAvailableCh)
	return r, wg
}

func teardown(testEnv *testEnv) {}
