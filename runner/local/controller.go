package local

import (
	"github.com/scootdev/scoot/runner"
)

func NewControllerAndStatuserRunner(cont runner.Controller, stat runner.Statuser) *ControllerAndStatuserRunner {
	return &ControllerAndStatuserRunner{cont, stat}
}

type ControllerAndStatuserRunner struct {
	cont runner.Controller
	stat runner.Statuser
}

func (r *ControllerAndStatuserRunner) Run(cmd *runner.Command) (runner.ProcessStatus, error) {
	return r.cont.Run(cmd)
}

func (r *ControllerAndStatuserRunner) Abort(run runner.RunId) (runner.ProcessStatus, error) {
	return r.cont.Abort(run)
}

func (r *ControllerAndStatuserRunner) StatusQuery(q runner.StatusQuery, opts runner.PollOpts) ([]runner.ProcessStatus, error) {
	return r.stat.StatusQuery(q, opts)
}

func (r *ControllerAndStatuserRunner) StatusQuerySingle(q runner.StatusQuery, opts runner.PollOpts) (runner.ProcessStatus, error) {
	return r.stat.StatusQuerySingle(q, opts)
}

func (r *ControllerAndStatuserRunner) Status(run runner.RunId) (runner.ProcessStatus, error) {
	return r.stat.Status(run)
}

func (r *ControllerAndStatuserRunner) StatusAll() ([]runner.ProcessStatus, error) {
	return r.stat.StatusAll()
}

func (r *ControllerAndStatuserRunner) Erase(run runner.RunId) error {
	return r.stat.Erase(run)
}
