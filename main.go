package main

import (
	"fmt"

	"github.com/bbuck/go-flag"
	"github.com/tree-server/trees/config"
	"github.com/tree-server/trees/errors"
	"github.com/tree-server/trees/log"
)

func main() {
	appLogger := log.Make("system", log.GetFileTarget(":stdout:"))

	init, err := flag.Bool("init, i", false, "Initialize TreeServer.toml with default values.")
	if err != nil {
		appLogger.Fatal(err, errors.ErrFailedToStart)
	}

	flag.Parse()

	if *init {
		config.LoadOrCreate()

		return
	}

	fmt.Println("Rest of program")
}
