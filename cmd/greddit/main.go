package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"encoding/json"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var credentials reddit.Credentials
	json.Unmarshal(byteValue, &credentials)

	client, err := reddit.NewClient(credentials)
	if err != nil {
		return
	}

	user, _, err := client.Account.Info(ctx)
	if err != nil {
		return
	}

	fmt.Println("Authenticated as ", user.Name)

	// with open('subs.txt') as f:
	//     lines = f.readlines()

	//     for subreddit_name in lines:
	//         subreddit = reddit.subreddit(subreddit_name)
	//         subreddit.subscribe()

	return
}
