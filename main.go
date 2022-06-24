package main

import (
	"os"

	"github.com/4nkitd/gobackup/config"
	"github.com/4nkitd/gobackup/model"
	"github.com/urfave/cli"
)

const (
	usage = "Easy full stack backup operations on UNIX-like systems"
)

var (
	modelName  = ""
	configFile = ""
	version    = "master"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "gobackup"
	app.Usage = usage

	app.Commands = []cli.Command{
		{
			Name: "perform",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "model, m",
					Usage:       "Model name that you want execute",
					Destination: &modelName,
				},
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Special a config file",
					Destination: &configFile,
				},
			},
			Action: func(c *cli.Context) error {
				config.Init(configFile)

				if len(modelName) == 0 {
					performAll()
				} else {
					performOne(modelName)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func performAll() {
	for _, modelConfig := range config.Models {
		m := model.Model{
			Config: modelConfig,
		}
		m.Perform()
	}
}

func performOne(modelName string) {
	for _, modelConfig := range config.Models {
		if modelConfig.Name == modelName {
			m := model.Model{
				Config: modelConfig,
			}
			m.Perform()
			return
		}
	}
}
