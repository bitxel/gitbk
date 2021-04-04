package gitbk

import (
	"github.com/bitxel/gitbk/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type Project struct {
	config config.Project
	auth transport.AuthMethod
}

func (p Project) Sync() error {
	startAt := time.Now()
	defer func() {
		log.Info().Str("project", p.config.WorkDir).Str("elapsed", time.Since(startAt).String()).Msg("sync done")
	}()
	repo, err := git.PlainOpen(p.config.WorkDir)
	if err != nil {
		return errors.Wrap(err, "open repo err")
	}
	wt, err := repo.Worktree()
	if err != nil {
		return errors.Wrap(err, "get worktree err")
	}
	status, err := wt.Status()
	if err != nil {
		return errors.Wrap(err, "get status err")
	}
	if len(status) == 0 {
		log.Info().Str("project", p.config.WorkDir).Msg("no changes found")
		return nil
	}

	_, err = wt.Add(".")
	if err != nil {
		return errors.Wrap(err, "add changes failed")
	}
	hash, err := wt.Commit("auto commit", &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name: "test",
			Email: "test@test.com",
			When: time.Now(),
		},
	})
	if err != nil {
		return errors.Wrap(err, "commit err")
	}
	log.Info().Str("project", p.config.WorkDir).Str("commit hash", hash.String())
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: p.auth,
	})
	if err != nil {
		return errors.Wrap(err, "push to remote error")
	}
	return nil
}

func NewProject(projectCfg config.Project) (*Project, error) {
	auth, err := PublicKeys(config.C.Global.Auth.PemFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "init project err")
	}
	return &Project{config: projectCfg, auth: auth}, nil
}

func SyncAll() (errs []error) {
	for _, projectCfg := range config.C.Project {
		Project, err := NewProject(projectCfg)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if err := Project.Sync(); err != nil {
			errs = append(errs, err)
		}
	}
	return
}

