package subreddits

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintSubcriptions(client *reddit.Client) (err error) {

	subs, _, err := client.Subreddit.Subscribed(ctx, nil)
	if err != nil {
		fmt.Println("Failed to retrieve subreddit list:", err)
		return
	}

	fmt.Printf("You are subscribed to %v subreddits:\n", len(subs))

	for _, s := range subs {
		fmt.Println(s.Name)
	}

	return nil
}
