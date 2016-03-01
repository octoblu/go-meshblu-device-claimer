package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("meshblu-device-claimer:main")

func main() {
	app := cli.NewApp()
	app.Name = "meshblu-device-claimer"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "path, p",
			EnvVar: "MESHBLU_JSON_PATH",
			Usage:  "Path to meshblu.json",
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	path := getOpts(context)
	logMsg := fmt.Sprintf("Path is %v", path)
	color.Green(logMsg)
}

func getOpts(context *cli.Context) string {
	path := context.String("path")

	if path == "" {
		cli.ShowAppHelp(context)

		if path == "" {
			color.Red("  Missing required flag --path or MESHBLU_JSON_PATH")
		}
		os.Exit(1)
	}

	return path
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
