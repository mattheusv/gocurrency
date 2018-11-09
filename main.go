package main

import (
	"github.com/msalcantara/gocurrency/api"
	"github.com/msalcantara/gocurrency/config"
)

func main() {
	config.Load()
	api.StartServer()
}
