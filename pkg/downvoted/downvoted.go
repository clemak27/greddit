package downvoted

import (
	"context"
	"fmt"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintDownvoted(rc client_wrapper.ClientWrapper) (err error) {

	downvoted, _ := GetDownvoted(rc)

	fmt.Printf("You have downvoted %v posts!\n", len(downvoted))

	for _, s := range downvoted {
		fmt.Println(s.Title)
	}

	return nil
}

func GetDownvoted(rc client_wrapper.ClientWrapper) (l []*reddit.Post, err error) {

	opts := reddit.ListOptions{Limit: 100}

	downvoted, _, err := rc.Downvoted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	if len(downvoted) == 100 {
		downvoted = append(downvoted, retrieveMore(downvoted, rc)...)
	}

	return downvoted, nil
}

func retrieveMore(subs []*reddit.Post, rc client_wrapper.ClientWrapper) []*reddit.Post {
	fli := subs[len(subs)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	nsl, _, err := rc.Downvoted(ctx, &reddit.ListUserOverviewOptions{
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
