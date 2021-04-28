package upvoted

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintUpvoted(client *reddit.Client) (err error) {

	// TODO implement
	// opts := reddit.ListOptions{Limit: 100}

	// subs, _, err := client.Subreddit.Subscribed(ctx, &reddit.ListSubredditOptions{
	// 	ListOptions: opts,
	// })

	// if err != nil {
	// 	fmt.Println("Failed to retrieve subreddit list:", err)
	// 	return
	// }

	// fmt.Printf("You are subscribed to %v subreddits:\n", len(subs))

	// for _, s := range subs {
	// 	fmt.Println(s.Name)
	// }

	return nil
}
