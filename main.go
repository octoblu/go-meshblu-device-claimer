package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	"github.com/octoblu/go-meshblu-device-claimer/meshblu"
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
			Name:   "claim-uri, c",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_CLAIM_URI",
			Usage:  "Base url to claim the meshblu device with, it will append /:uuid/:token.",
			Value:  "https://app.octoblu.com/node-wizard/claim",
		},
		cli.StringFlag{
			Name:   "type, t",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_TYPE",
			Usage:  "Device type to register",
		},
		cli.StringFlag{
			Name:   "meshblu-uri, m",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_MESHBLU_URI",
			Usage:  "Meshblu server to register the device with",
			Value:  "https://meshblu.octoblu.com:443",
		},
		cli.StringFlag{
			Name:   "path, p",
			EnvVar: "MESHBLU_DEVICE_CLAIMER_PATH",
			Usage:  "Path to meshblu.json",
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	claimURI, deviceType, meshbluURI, path := getOpts(context)
	fmt.Println("getOpts", claimURI, deviceType, meshbluURI, path)

	client := meshblu.New(meshbluURI)
	config, err := client.Register(deviceType)
	if err != nil {
		log.Fatalln("Error on client register:", err.Error())
	}

	configJSON, err := config.ToJSON()
	if err != nil {
		log.Fatalln("Error on toJSONing:", err.Error())
	}

	ioutil.WriteFile(path, configJSON, 0644)

	fmt.Println("config", config)
}

func getOpts(context *cli.Context) (string, string, string, string) {
	claimURI := context.String("claim-uri")
	deviceType := context.String("type")
	meshbluURI := context.String("meshblu-uri")
	path := context.String("path")

	if path == "" {
		cli.ShowAppHelp(context)

		if path == "" {
			color.Red("  Missing required flag --path, -p, or MESHBLU_DEVICE_CLAIMER_PATH")
		}

		os.Exit(1)
	}

	return claimURI, deviceType, meshbluURI, path
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
