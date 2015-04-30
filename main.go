package main

import (
	"fmt"

	"github.com/bbuck/go-flag"
	"github.com/seer-server/seer/config"
	"github.com/seer-server/seer/errors"
	"github.com/seer-server/seer/log"
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
