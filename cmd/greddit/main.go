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
	"gitlab.com/clemak27/greddit/pkg/subreddits"
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
			{
				Name:    "subreddits",
				Usage:   "interact with subreddits",
				Aliases: []string{"sr"},
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all subreddits you are subscribed to",
						Action: func(c *cli.Context) error {
							credentials, _ := getConfig(configPath)
							client, _ := authentication.GetClient(credentials)
							subreddits.PrintSubcriptions(client)
							return nil
						},
					},
					{
						Name:  "subscribe",
						Usage: "subscribe to subreddit with `NAME`",
						Action: func(c *cli.Context) error {
							if c.Args().Len() == 0 {
								fmt.Println("missing argument")
								return nil
							}
							fmt.Println(c.Args().Get(0))
							return nil
						},
					},
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
