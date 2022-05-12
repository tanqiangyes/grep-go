// Package grep
// @Author: tanqiangyes
// @Emainl: 826285820@qq.com
// @Date: 2022-05-11 21:09:05
// @Description: a grep written in go. just for study linux and go.
package cmd

import (
	"fmt"
	"github.com/tanqiangyes/grep-go/in_errors"
	"github.com/tanqiangyes/grep-go/reader"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Version : a version for command
const Version = "0.0.1"

func Main() {
	app := &cli.App{
		Name:                 "grep-go",
		Usage:                "a grep written in go. just for study linux and go.",
		Version:              Version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			// a flag for regexp
			&cli.BoolFlag{
				Name:     "regexp",
				Aliases:  []string{"e"},
				Usage:    "This flag treats patterns as regular expressions and looks for content in the corresponding file that matches the regular expression.",
				Required: false,
				Hidden:   false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     "recursive",
				Aliases:  []string{"r"},
				Usage:    "Whether to look recursively in the path.",
				Required: false,
				Hidden:   false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     "ignore-case",
				Aliases:  []string{"i"},
				Usage:    "Ignore  case  distinctions,  so that characters that differ, only in case match each other.",
				Required: false,
				Hidden:   false,
				Value:    false,
			},
		},
		Action: func(c *cli.Context) error {
			//todo

			//fmt.Println(c.Bool("r"), c.Bool("e"), c.Bool("i"))
			//return nil

			//return nil
			var (
				read reader.Reader
				err  error
			)
			//
			// if there only one argument, show help for the user.
			if c.Args().Len() < 1 {
				return in_errors.ErrArgs
			}

			// if no file, so args len is 1, we can try read data from  stdin.
			if c.Args().Len() == 1 {
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
					log.Fatal("The command is intended to work with pipes.")
					return nil
				}
				// readCloser is stdin
				read, err = reader.NewStdReader(os.Stdin, c.Args().First(), c.Bool("e"), c.Bool("i"))
				if err != nil {
					return err
				}
			} else {
				path := c.Args().Slice()[1:]
				fmt.Println(path)
				// we should open files or dir, and then read from it.
				read, err = reader.NewMultiReader(path, c.Args().First(), c.Bool("r"), c.Bool("e"), c.Bool("i"))
				if err != nil {
					return err
				}
			}
			read.Run()
			read.Print()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GetStdinOrFileData() {

}
