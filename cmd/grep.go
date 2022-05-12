// Package grep
// @Author: tanqiangyes
// @Emainl: 826285820@qq.com
// @Date: 2022-05-11 21:09:05
// @Description: a grep written in go. just for study linux and go.
package cmd

import (
	"fmt"
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
				Aliases:  []string{"E"},
				Usage:    "This flag treats patterns as regular expressions and looks for content in the corresponding file that matches the regular expression.",
				Required: false,
				Hidden:   false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     "fixed-strings",
				Aliases:  []string{"F"},
				Usage:    "This flag treats patterns as fixed strings and looks for the same content in the corresponding file.",
				Required: false,
				Hidden:   false,
				Value:    false,
			},
		},
		Action: func(c *cli.Context) error {
			//todo
			fmt.Println(c.Args(), os.Args)
			//return nil
			//var (
			//	r   io.ReadCloser = os.Stdin
			//	err error
			//)
			//
			//// if there only one argument, show help for the user.
			//if c.Args().Len() < 1 {
			//	return fmt.Errorf("error: too few arguments, please --help  to show help text.")
			//}
			//
			//// if no file, so args len is 1, we can try read data from  stdin.
			//if c.Args().Len() == 1 {
			//	stat, _ := os.Stdin.Stat()
			//	if (stat.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
			//		log.Fatal("The command is intended to work with pipes.")
			//		return nil
			//	}
			//	r = os.Stdin
			//} else {
			//
			//}

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
