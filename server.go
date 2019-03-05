package main

import "github.com/labstack/gommon/log"

const (
	DefaultConfFilePath = "conf/conf.toml"
)

var (
	confFilePath string
)

func main() {
	log.Debug("run with conf: %s", confFilePath)
	router.RunSubdomains(confFilePath)
}
