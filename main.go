package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/msalcantara/gocurrency/api"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var port int
	a := kingpin.New(filepath.Base(os.Args[0]), "web scraping to provide currency quote data in REST api")
	a.Arg("port", "http port").Envar("PORT").Default("8080").IntVar(&port)
	if _, err := a.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error to start api: %v\n", err)
		os.Exit(1)
	}
	server := api.NewServer(port)
	if err := server.Run(); err != nil {
		fmt.Printf("Error to start api: %v\n", err)
		os.Exit(1)
	}
}
