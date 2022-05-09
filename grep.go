package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "lotus-miner",
		Usage:                "Filecoin decentralized storage network miner",
		Version:              "0.0.1",
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			//todo
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
