package saved

import (
	"context"
	"fmt"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintSaved(rc client_wrapper.ClientWrapper) (err error) {

	saved, _ := GetSaved(rc)

	fmt.Printf("You have saved %v posts!\n", len(saved))

	for _, s := range saved {
		fmt.Println(s.Title)
	}

	return nil
}

func GetSaved(rc client_wrapper.ClientWrapper) (l []*reddit.Post, err error) {

	opts := reddit.ListOptions{Limit: 100}

	saved, _, _, err := rc.Saved(ctx, &reddit.ListUserOverviewOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve post list:", err)
		return
	}

	if len(saved) == 100 {
		saved = append(saved, retrieveMore(saved, rc)...)
	}

	return saved, nil
}

func retrieveMore(subs []*reddit.Post, rc client_wrapper.ClientWrapper) []*reddit.Post {
	fli := subs[len(subs)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	nsl, _, _, err := rc.Saved(ctx, &reddit.ListUserOverviewOptions{
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
