package config

import (
	"encoding/json"
	"time"
)

var C = Config{
	Global: Global{
		BackupInterval: time.Duration(time.Hour),
		Message: "auto commit by gitbk",
	},
}

type Config struct {
	Global Global
	Project []Project
}

func (c Config) String() string {
	j, _ := json.Marshal(c)
	return string(j)
}

type Global struct {
	BackupInterval time.Duration `mapstructure:"backup_interval"`
	LogLevel int8
	Test int
	Auth Auth
	Message string
}

const (
	AuthMethodPublicKeys = "public_keys"
)

type Auth struct {
	Method string
	PemFilePath string `mapstructure:"pem_file_path"`
	PemFilePassword string `mapstructure:"pem_file_password"`
}

type Project struct {
	URL string
	Branch string
	Message string
	WorkDir string
}