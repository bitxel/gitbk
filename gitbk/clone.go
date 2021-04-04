package gitbk

import (
	"github.com/bitxel/gitbk/config"
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
)

func Clone(cfg config.Project) error {
	auth, err := PublicKeys(config.C.Global.Auth.PemFilePath)
	if err != nil {
		return errors.Wrap(err, "init project err")
	}

	_, err = git.PlainClone(cfg.WorkDir, false, &git.CloneOptions{
		URL: cfg.URL,
		Progress: os.Stdout,
		Auth: auth,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return errors.Wrap(err, "clone from remote failed")
	}
	return nil
}
