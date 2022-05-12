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
	"github.com/tanqiangyes/grep-go/tools"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Version : a version for command
const Version = "0.0.1"

func Main() {
	app := &cli.App{
		Name:                 "grep",
		Usage:                "a grep written in go. just for study linux and go.",
		Version:              Version,
		EnableBashCompletion: true,
		ArgsUsage: `grep PATTERNS [FILE...]
       grep -e  PATTERNS ... [FILE...]
       grep -f -r PATTERN_FILE ... [FILE...]
					`,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "tanqiangyes",
				Email: "826285820@qq.com",
			},
		},
		Flags: []cli.Flag{
			// a flag for regexp
			&cli.BoolFlag{
				Name:    "regexp",
				Aliases: []string{"e"},
				Usage:   "This flag treats patterns as regular expressions and looks for content in the corresponding file that matches the regular expression.",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "Whether to look recursively in the path.",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "get pattern from the file.",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "line-number",
				Aliases: []string{"n"},
				Usage:   "print the number of lines matched.",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "ignore-case",
				Aliases: []string{"i"},
				Usage:   "Ignore  case  distinctions,  so that characters that differ, only in case match each other.",
				Value:   false,
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
				PrintError(c, in_errors.ErrArgs)
				return nil
			}

			var finders []reader.Finder
			// pattern, a file or a string, if a file, we must read the file, then make slice of finders. In the file, we use "\n" as a separator.
			pattern := c.Args().First()
			//regexp? default false
			regexp := c.Bool("e")
			// ignore case? default false
			caseIgnore := c.Bool("i")
			if c.Bool("f") {
				//this is a file, so we should read from it, then make all finder.
				patterns, err := tools.ReadFile(pattern)
				if err != nil {
					PrintError(c, err)
					return nil
				}
				for _, p := range patterns {
					finder, err := reader.NewFinder(p, regexp, caseIgnore)
					if err != nil {
						PrintError(c, err)
						return nil
					}
					finders = append(finders, finder)
				}
			} else {
				// only one pattern, so we make a finder.
				finder, err := reader.NewFinder(pattern, regexp, caseIgnore)
				if err != nil {
					PrintError(c, err)
					return nil
				}
				finders = append(finders, finder)
			}

			// if no file, so args len is 1, we can try read data from  stdin.
			if c.Args().Len() == 1 {
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
					PrintError(c, in_errors.ErrMustWorkWithPipes)
					return nil
				}
				// readCloser is stdin, from the code, we can not deal err, because no error here.
				read, _ = reader.NewStdReader(os.Stdin, finders)
			} else {
				path := c.Args().Slice()[1:]
				fmt.Println(path)
				// we should open files or dir, and then read from it.
				read, err = reader.NewMultiReader(path, finders, c.Bool("r"))
				if err != nil {
					PrintError(c, err)
					return nil
				}
			}
			read.Run()
			PrintMatch(c, read.Result())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintError(cctx *cli.Context, err error) {
	fmt.Println(err.Error())
	fmt.Println()
	// no need to deal with error
	cli.ShowAppHelp(cctx)
}

func PrintMatch(cctx *cli.Context, result []reader.MatchRes) {
	format := ""
	line := cctx.Bool("n")
	for _, res := range result {
		lens := len(res.Lines)
		if lens < 1 {
			continue
		}
		format += fmt.Sprintf("File: %v\n", res.Filename)
		for i := 0; i < lens; i++ {
			if line {
				format += fmt.Sprintf("line %v:%v\n", res.Lines, res.MatchString)
			} else {
				format += fmt.Sprintf("%v\n", res.MatchString)
			}
		}
		format += fmt.Sprintf("\n\n")
	}
	if format == "" {
		fmt.Println("No matching rows")
		return
	}
	fmt.Println(format)
}
