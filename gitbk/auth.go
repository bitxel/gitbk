package gitbk

import (
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"log"
	"os"
)

var (
	defaultPrivateKeyFile = ""
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to get user's home folder, err:%s", err)
	}
	defaultPrivateKeyFile = homeDir + "/.ssh/id_rsa"
}

func PublicKeys(path string) (*ssh.PublicKeys, error) {
	if path == "" {
		path = defaultPrivateKeyFile
	}
	path, err := homedir.Expand(path)
	if err != nil {
		return nil, errors.Wrap(err, "expand path err")
	}
	_, err = os.Stat(path)
	if err != nil {
		return nil, errors.Wrap(err, "private key file not found")
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", path, "")
	if err != nil {
		return nil, errors.Wrap(err, "read private key file failed")
	}
	return publicKeys, nil
}

