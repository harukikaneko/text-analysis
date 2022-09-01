package main

import (
	"echo-get-started/config"
	"echo-get-started/rest/server"
)

func main() {
	config.InitConfig()
	server.Run()
}
