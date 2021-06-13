package upvoted

import (
	"context"
	"fmt"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintUpvoted(client *reddit.Client) (err error) {

	upvoted, _ := GetUpvoted(&client_wrapper.RedditClient{Client: client})

	fmt.Printf("You have upvoted %v posts!\n", len(upvoted))

	for _, s := range upvoted {
		fmt.Println(s.Title)
	}

	return nil
}

func GetUpvoted(rc client_wrapper.ClientFunctions) (l []*reddit.Post, err error) {

	opts := reddit.ListOptions{Limit: 100}

	upvoted, _, err := rc.Upvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	if len(upvoted) == 100 {
		upvoted = append(upvoted, retrieveMore(upvoted, rc)...)
	}

	return upvoted, nil
}

func retrieveMore(subs []*reddit.Post, rc client_wrapper.ClientFunctions) []*reddit.Post {
	fli := subs[len(subs)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	nsl, _, err := rc.Upvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: nopts,
	})
	if err != nil {
		fmt.Println("Failed to retrieve subreddit list:", err)
	}
	if len(nsl) == 100 {
		nsl = append(nsl, retrieveMore(nsl, rc)...)
	}
	return nsl
}
