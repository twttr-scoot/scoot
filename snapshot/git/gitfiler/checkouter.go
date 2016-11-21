package gitfiler

import (
	"fmt"
	"os/exec"

	"github.com/scootdev/scoot/snapshot"
	"github.com/scootdev/scoot/snapshot/git/repo"
)

func NewCheckouter(repos *RepoPool) *Checkouter {
	return &Checkouter{repos: repos}
}

// Checkouter checks out by checking out in a repo from pool
type Checkouter struct {
	repos *RepoPool
}

// An arbitrary revision. As mentioned below, we should get rid of this altogether
const DEFAULT_REV = "1dda9fbde682e4922a0d5709c5539f573db4cc54"

// Checkout checks out id (a raw git sha) into a Checkout.
// It does this by making a new clone (via reference) and checking out id.
// TODO(dbentley): if id is not found because it is present in upstream repos but not here,
// we should fetch it and check it out.
func (c *Checkouter) Checkout(id string) (co snapshot.Checkout, err error) {
	repo, repoErr := c.repos.Get()
	if repoErr != nil {
		return nil, repoErr
	}

	// release if we aren't using it
	defer func() {
		if err != nil || recover() != nil {
			c.repos.Release(repo, repoErr)
		}
	}()

	// TODO(dbentley): do more ot validate id. E.g., don't let "HEAD" or "master" slip through
	if id == "" {
		id = DEFAULT_REV
	}

	// -d removes directories. -x ignores gitignore and removes everything.
	// -f is force. -f the second time removes directories even if they're git repos themselves
	cmds := [][]string{
		{"clean", "-f", "-f", "-d", "-x"},
		{"checkout", id},
	}

	// fetchCmds  := [][]string{
	// 	{"fetch"},
	// 	{"clean", "-f", "-f", "-d", "-x"},
	// 	{"checkout", id}
	// }

	if err := c.runGitCmds(cmds, repo); err != nil {
		// try fetching for new commits
		err = c.runGitCmds(append([][]string{{"fetch"}}, cmds...), repo)
		if err != nil {
			return nil, err
		}
	}

	// for _, argv := range cmds {
	// 	if _, err := repo.Run(argv...); err != nil {
	// 		return nil, fmt.Errorf("gitfiler.Checkouter.Checkout: %v", err)
	// 	}
	// }
	return &Checkout{repo: repo, id: id, pool: c.repos}, nil
}

func (c *Checkouter) runGitCmds(cmds [][]string, repo *repo.Repository) error {
	for _, argv := range cmds {
		if _, err := repo.Run(argv...); err != nil {
			return fmt.Errorf("Unable to run git commands: %v", err)
		}
	}
	return nil
}

func (c *Checkouter) CheckoutAt(id string, dir string) (co snapshot.Checkout, err error) {
	co, err = c.Checkout(id)
	if err != nil {
		return nil, err
	}
	defer co.Release()

	cmd := exec.Command("sh", "-c", fmt.Sprintf("cp -r %s/* %s", co.Path(), dir))
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return &UnmanagedCheckout{id: id, dir: dir}, nil
}

// Implement noop ingest so this Checkouter can be passed around as a Filer.
func (c *Checkouter) Ingest(string) (string, error)               { return "", nil }
func (c *Checkouter) IngestMap(map[string]string) (string, error) { return "", nil }
func (c *Checkouter) AsFiler() snapshot.Filer                     { return c }

// Checkout holds one repo that is checked out to a specific ID
type Checkout struct {
	repo *repo.Repository
	id   string
	pool *RepoPool
}

func (c *Checkout) Path() string {
	return c.repo.Dir()
}

func (c *Checkout) ID() string {
	return c.id
}

func (c *Checkout) Release() error {
	if c.pool != nil {
		c.pool.Release(c.repo, nil)
		c.pool = nil
	}
	return nil
}

// User-owned checkout.
type UnmanagedCheckout struct {
	id  string
	dir string
}

func (c *UnmanagedCheckout) Path() string {
	return c.dir
}

func (c *UnmanagedCheckout) ID() string {
	return c.id
}

func (c *UnmanagedCheckout) Release() error {
	return nil
}
