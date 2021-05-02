package upvoted

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintUpvoted(client *reddit.Client) (err error) {

	opts := reddit.ListOptions{Limit: 100}

	upvoted, _, err := client.User.Upvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	fmt.Printf("You have upvoted %v posts!\n", len(upvoted))

	for _, s := range upvoted {
		fmt.Println(s.Title)
	}

	return nil
}

func GetUpvoted(client *reddit.Client) (l []*reddit.Post) {

	opts := reddit.ListOptions{Limit: 100}

	upvoted, _, err := client.User.Upvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})
	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	return upvoted
}
