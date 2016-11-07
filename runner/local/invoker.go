package local

import (
	"fmt"
	"log"
	"time"

	"github.com/scootdev/scoot/runner"
	"github.com/scootdev/scoot/runner/execer"
	"github.com/scootdev/scoot/snapshot"
)

func NewInvoker(exec execer.Execer, checkouter snapshot.Checkouter, outputCreator runner.OutputCreator) *Invoker {
	return &Invoker{exec, checkouter, outputCreator}
}

type Invoker struct {
	execer        execer.Execer
	checkouter    snapshot.Checkouter
	outputCreator runner.OutputCreator
}

type checkoutAndError struct {
	checkout snapshot.Checkout
	err      error
}

func (in *Invoker) Run(cmd *runner.Command, id runner.RunId, abortCh chan struct{}, updateCh chan runner.ProcessStatus) runner.ProcessStatus {
	log.Printf("runner/local/invoker.go: Run id %v with cmd %+v", id, cmd)

	var timeoutCh <-chan time.Time
	if cmd.Timeout > 0 { // Timeout if applicable
		timeout := time.NewTicker(cmd.Timeout)
		timeoutCh = timeout.C
		defer timeout.Stop()
	}

	checkoutCh := make(chan checkoutAndError)
	var checkout snapshot.Checkout
	var err error
	go func() {
		checkout, err = in.checkouter.Checkout(cmd.SnapshotId)
		checkoutCh <- checkoutAndError{checkout, err}
	}()

	drainCheckout := func() {
		checkoutAndErr := <-checkoutCh
		if checkoutAndErr.err != nil {
			return
		}
		checkoutAndErr.checkout.Release()
	}

	select {
	case <-abortCh:
		go drainCheckout()
		return runner.AbortStatus(id)
	case <-timeoutCh:
		go drainCheckout()
		return runner.TimeoutStatus(id)
	case checkoutAndErr := <-checkoutCh:
		checkout, err = checkoutAndErr.checkout, checkoutAndErr.err
	}

	defer checkout.Release()

	log.Printf("runner.local.invoker.Run cmd: %+v checkout: %v", cmd, checkout.Path())

	stdout, err := in.outputCreator.Create(fmt.Sprintf("%s-stdout", id))
	if err != nil {
		return runner.ErrorStatus(id, fmt.Errorf("could not create stdout: %v", err))
	}
	defer stdout.Close()
	stderr, err := in.outputCreator.Create(fmt.Sprintf("%s-stderr", id))
	if err != nil {
		return runner.ErrorStatus(id, fmt.Errorf("could not create stderr: %v", err))
	}
	defer stderr.Close()

	p, err := in.execer.Exec(execer.Command{
		Argv:   cmd.Argv,
		Dir:    checkout.Path(),
		Stdout: stdout,
		Stderr: stderr,
	})
	if err != nil {
		return runner.ErrorStatus(id, fmt.Errorf("could not exec: %v", err))
	}

	updateCh <- runner.RunningStatus(id, stdout.URI(), stderr.URI())

	processCh := make(chan execer.ProcessStatus, 1)
	go func() { processCh <- p.Wait() }()
	var st execer.ProcessStatus

	// Wait for process to complete (or cancel if we're told to)
	select {
	case <-abortCh:
		p.Abort()
		return runner.AbortStatus(id)
	case <-timeoutCh:
		p.Abort()
		return runner.TimeoutStatus(id)
	case st = <-processCh:
	}

	switch st.State {
	case execer.COMPLETE:
		return runner.CompleteStatus(id, st.ExitCode)
	case execer.FAILED:
		return runner.ErrorStatus(id, fmt.Errorf("error execing: %v", st.Error))
	default:
		return runner.ErrorStatus(id, fmt.Errorf("unexpected exec state: %v", st.State))
	}
}
