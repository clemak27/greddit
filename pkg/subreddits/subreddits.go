package subreddits

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintSubcriptions(rc client_wrapper.ClientWrapper) (err error) {

	subs, _ := GetSubscriptions(rc)

	fmt.Printf("You are subscribed to %v subreddits:\n", len(subs))

	for _, s := range subs {
		fmt.Println(s.Name)
	}

	return nil
}

func GetSubscriptions(rc client_wrapper.ClientWrapper) (l []*reddit.Subreddit, err error) {

	opts := reddit.ListOptions{Limit: 100}

	subs, _, err := rc.Subscribed(ctx, &reddit.ListSubredditOptions{
		ListOptions: opts,
	})

	if err != nil {
		fmt.Println("Failed to retrieve subreddit list:", err)
		return
	}

	if len(subs) == 100 {
		subs = append(subs, retrieveMore(subs, rc)...)
	}

	return subs, nil
}

func Subscribe(rc client_wrapper.ClientWrapper, name string) (err error) {

	rc.Subscribe(ctx, name)
	fmt.Printf("Subscribed to %v\n", name)

	return nil
}

func SubscribeFromFile(rc client_wrapper.ClientWrapper, subPath string) (err error) {

	subredditNames, _ := scanLines(subPath)

	for _, v := range subredditNames {
		rc.Subscribe(ctx, v)
		fmt.Printf("Subscribed to %v\n", v)
	}

	return nil
}

func Unsubscribe(rc client_wrapper.ClientWrapper, subredditName string) (err error) {

	rc.Unsubscribe(ctx, subredditName)
	fmt.Printf("Unsubscribed from %v\n", subredditName)

	return nil
}

func scanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func retrieveMore(subs []*reddit.Subreddit, rc client_wrapper.ClientWrapper) []*reddit.Subreddit {
	fli := subs[len(subs)-1].FullID
	nopts := reddit.ListOptions{Limit: 100, After: fli}
	nsl, _, err := rc.Subscribed(ctx, &reddit.ListSubredditOptions{
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
