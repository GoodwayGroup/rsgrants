package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/GoodwayGroup/rsgrants/info"
	"github.com/clok/cdocs"
	"github.com/urfave/cli/v2"
)

var version string

func main() {
	// Generate the install-manpage command
	im, err := cdocs.InstallManpageCommand(&cdocs.InstallManpageCommandInput{
		AppName: info.AppName,
	})
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:     info.AppName,
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Derek Smith",
				Email: "dsmith@goodwaygroup.com",
			},
			{
				Name: info.AppRepoOwner,
			},
		},
		Copyright:            "(c) 2021 Goodway Group",
		HelpName:             info.AppName,
		Usage:                "AWS Redshift grants tool",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			im,
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print version info",
				Action: func(c *cli.Context) error {
					fmt.Printf("%s %s (%s/%s)\n", info.AppName, version, runtime.GOOS, runtime.GOARCH)
					return nil
				},
			},
		},
	}

	if os.Getenv("DOCS_MD") != "" {
		docs, err := cdocs.ToMarkdown(app)
		if err != nil {
			panic(err)
		}
		fmt.Println(docs)
		return
	}

	if os.Getenv("DOCS_MAN") != "" {
		docs, err := cdocs.ToMan(app)
		if err != nil {
			panic(err)
		}
		fmt.Println(docs)
		return
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}