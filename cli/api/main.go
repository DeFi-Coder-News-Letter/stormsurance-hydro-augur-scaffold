package main

import (
	_ "github.com/joho/godotenv/autoload"
)

import (
	"context"
	"github.com/HydroProtocol/hydro-box-augur/api"
	"github.com/HydroProtocol/hydro-box-augur/cli"
	"os"
)

func run() int {
	ctx, stop := context.WithCancel(context.Background())

	go cli.WaitExitSignal(stop)
	api.StartServer(ctx)

	return 0
}

func main() {
	os.Exit(run())
}
