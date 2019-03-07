package main

import (
	"conf"
	"router"
)

var (
	defaultConfigFile = "src/conf/conf.toml"
)

func main() {
	err := conf.InitConfig(defaultConfigFile)
	//println(err.Error())
	if err != nil {
		println(err)
	} else {
		_ = router.InitRouters()
	}
}