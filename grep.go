// Package main
// @Author: tanqiangyes
// @Emainl: 826285820@qq.com
// @Date: 2022-05-11 21:09:05
// @Description: a grep written in go. just for study linux and go.
package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Version : a version for command
const Version = "0.0.1"

func main() {
	app := &cli.App{
		Name:                 "lotus-miner",
		Usage:                "Filecoin decentralized storage network miner",
		Version:              Version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "",
				Aliases:     nil,
				Usage:       "",
				EnvVars:     nil,
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				Value:       false,
				DefaultText: "",
				Destination: nil,
				HasBeenSet:  false,
			},
		},
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
