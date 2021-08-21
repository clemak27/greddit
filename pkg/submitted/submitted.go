package submitted

import (
	"context"
	"fmt"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintSubmitted(rc client_wrapper.ClientWrapper) (err error) {

	submitted, _ := GetSubmitted(rc)

	fmt.Printf("You have submitted %v posts!\n", len(submitted))

	for _, s := range submitted {
		fmt.Println(s.Title)
	}

	return nil
}

func GetSubmitted(rc client_wrapper.ClientWrapper) (l []*reddit.Post, err error) {

	opts := reddit.ListOptions{Limit: 100}

	submitted, _, err := rc.Submitted(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	if len(submitted) == 100 {
		submitted = append(submitted, retrieveMore(submitted, rc)...)
	}

	return submitted, nil
}

func retrieveMore(subs []*reddit.Post, rc client_wrapper.ClientWrapper) []*reddit.Post {
	fli := subs[len(subs)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	nsl, _, err := rc.Submitted(ctx, &reddit.ListUserOverviewOptions{
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
