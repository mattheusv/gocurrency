package main

import (
	"github.com/msalcantara/currency-quote-api/api"
	"github.com/msalcantara/currency-quote-api/config"
)

func main() {
	config.Load()
	api.StartServer()
}
