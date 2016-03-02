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
			EnvVar: "MESHBLU_DEVICE_CLAIMER_PATH",
			Usage:  "Path to meshblu.json",
		},
		cli.StringFlag{
			Name:   "meshblu-uri, m",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_MESHBLU_URI",
			Usage:  "Meshblu server to register the device with",
			Value:  "https://meshblu.octoblu.com:443",
		},
		cli.StringFlag{
			Name:   "claim-uri, c",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_CLAIM_URI",
			Usage:  "Base url to claim the meshblu device with, it will append /:uuid/:token.",
			Value:  "https://app.octoblu.com/node-wizard/claim",
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	path, meshbluURI, claimURI := getOpts(context)
	logMsg := fmt.Sprintf("Opts are %v %v %v", path, meshbluURI, claimURI)
	color.Green(logMsg)
}

func getOpts(context *cli.Context) (string, string, string) {
	path := context.String("path")
	meshbluURI := context.String("meshblu-uri")
	claimURI := context.String("claim-uri")

	if path == "" {
		cli.ShowAppHelp(context)

		if path == "" {
			color.Red("  Missing required flag --path, -p, or MESHBLU_DEVICE_CLAIMER_JSON_PATH")
		}

		os.Exit(1)
	}

	return path, meshbluURI, claimURI
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
