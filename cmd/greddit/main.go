package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"encoding/json"

	"github.com/clemak27/greddit/pkg/authentication"
	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/clemak27/greddit/pkg/downvoted"
	"github.com/clemak27/greddit/pkg/export"
	"github.com/clemak27/greddit/pkg/saved"
	"github.com/clemak27/greddit/pkg/submitted"
	"github.com/clemak27/greddit/pkg/subreddits"
	"github.com/clemak27/greddit/pkg/upvoted"
	"github.com/urfave/cli/v2"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func main() {
	var configPath string
	var subPath string
	var outputFormat string

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
					getClientWrapper(configPath)
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
							wrapper := getClientWrapper(configPath)
							subreddits.PrintSubcriptions(wrapper)
							return nil
						},
					},
					{
						Name:  "subscribe",
						Usage: "subscribe to subreddit with `NAME`",
						Action: func(c *cli.Context) error {

							wrapper := getClientWrapper(configPath)

							if subPath != "" {
								subreddits.SubscribeFromFile(wrapper, subPath)
								return nil
							}

							if !c.Args().Present() {
								fmt.Println("missing argument, specify a subreddit name")
								return nil
							}
							for _, v := range c.Args().Slice() {
								subreddits.Subscribe(wrapper, v)
							}
							return nil
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file",
								Aliases:     []string{"f"},
								Usage:       "File containing a list of subreddits to subscribe to, separated by newlines.",
								Destination: &subPath,
							},
						},
					},
					{
						Name:  "unsubscribe",
						Usage: "unsubscribe from subreddit with `NAME`",
						Action: func(c *cli.Context) error {

							wrapper := getClientWrapper(configPath)

							if !c.Args().Present() {
								fmt.Println("missing argument, specify a subreddit name")
								return nil
							}

							for _, v := range c.Args().Slice() {
								subreddits.Unsubscribe(wrapper, v)
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "upvoted",
				Usage: "upvoted posts",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all posts you have upvoted",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							upvoted.PrintUpvoted(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "saved",
				Usage: "saved posts",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all posts you have saved",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							saved.PrintSaved(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "saved-comments",
				Usage: "saved comments",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all comments you have saved",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							saved.PrintSavedComments(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "downvoted",
				Usage: "downvoted posts",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all posts you have downvoted",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							downvoted.PrintDownvoted(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "submitted",
				Usage: "submitted posts",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all posts you have submitted",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							submitted.PrintSubmitted(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "submitted-comments",
				Usage: "submitted comments",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "prints a list of all comments you have submitted",
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							submitted.PrintSubmittedComments(wrapper)
							return nil
						},
					},
				},
			},
			{
				Name:  "export",
				Usage: "export posts",
				Subcommands: []*cli.Command{
					{
						Name:  "upvoted",
						Usage: "exports a list of all posts you have upvoted",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Posts(wrapper, outputFormat, "upvoted")
							handleError(err)
							return nil
						},
					},
					{
						Name:  "saved",
						Usage: "exports a list of all posts you have saved",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Posts(wrapper, outputFormat, "saved")
							handleError(err)
							return nil
						},
					},
					{
						Name:  "saved-comments",
						Usage: "exports a list of all comments you have saved",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Comments(wrapper, outputFormat, "saved-comments")
							handleError(err)
							return nil
						},
					},
					{
						Name:  "downvoted",
						Usage: "exports a list of all posts you have downvoted",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Posts(wrapper, outputFormat, "downvoted")
							handleError(err)
							return nil
						},
					},
					{
						Name:  "submitted",
						Usage: "exports a list of all posts you have submitted",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Posts(wrapper, outputFormat, "submitted")
							handleError(err)
							return nil
						},
					},
					{
						Name:  "submitted-comments",
						Usage: "exports a list of all comments you have submitted",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "format",
								Aliases:     []string{"f"},
								Value:       "md",
								Usage:       "output format of the export",
								Destination: &outputFormat,
							}},
						Action: func(c *cli.Context) error {
							wrapper := getClientWrapper(configPath)
							err := export.Comments(wrapper, outputFormat, "submitted-comments")
							handleError(err)
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

func getClientWrapper(configPath string) client_wrapper.ClientWrapper {

	var credentials reddit.Credentials

	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &credentials)

	client, _ := authentication.GetClient(credentials)
	wrapper := client_wrapper.RedditClient{Client: client}

	return &wrapper
}

func handleError(err error) {
	if err != nil {
		fmt.Println("An error occured:")
		fmt.Println(err)
		os.Exit(1)
	}
}
