package upvoted

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintUpvoted(client *reddit.Client) (err error) {

	upvoted := GetUpvoted(client)

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

	if len(upvoted) == 100 {
		upvoted = append(upvoted, retrieveMore(upvoted, client)...)
	}

	return upvoted
}

func retrieveMore(posts []*reddit.Post, client *reddit.Client) []*reddit.Post {
	fli := posts[len(posts)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	npl, _, err := client.User.Upvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: nopts,
	})
	if err != nil {
		fmt.Println("Failed to retrieve subreddit list:", err)
	}
	if len(npl) == 100 {
		npl = append(npl, retrieveMore(npl, client)...)
	}
	return npl
}
