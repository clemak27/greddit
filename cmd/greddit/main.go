package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"encoding/json"

	"github.com/urfave/cli/v2"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"gitlab.com/clemak27/greddit/pkg/authentication"
)

var ctx = context.Background()

func main() {
	var configPath string

	app := &cli.App{
		Name:  "greddit",
		Usage: "greddit is a cli utility to interact with the reddit api.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `FILE`. Should contain authentication info.",
				Value:       "./config.json",
				Destination: &configPath,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "authenticate",
				Usage: "authenticates with the reddit api. This command is mainly for testing if the config is set correctly.",
				Action: func(c *cli.Context) error {
					credentials, _ := getConfig(configPath)
					authentication.GetClient(credentials)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig(path string) (credentials reddit.Credentials, err error) {

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &credentials)

	return credentials, nil
}
