package main

import (
	"fmt"
	"gowebstarter/adapter"
	"gowebstarter/infra"
	"gowebstarter/infra/config"
)

func main() {
	infra.Init()
	r := adapter.InitGin()
	address := fmt.Sprintf("%s:%s", config.GetConf().Server.Host, config.GetConf().Server.Port)
	r.Run(address)
}
