package subreddits

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func PrintSubcriptions(client *reddit.Client) (err error) {

	opts := reddit.ListOptions{Limit: 100}

	subs, _, err := client.Subreddit.Subscribed(ctx, &reddit.ListSubredditOptions{
		ListOptions: opts,
	})

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

func Subscribe(client *reddit.Client, name string) (err error) {

	client.Subreddit.Subscribe(ctx, name)
	fmt.Printf("Subscribed to %v\n", name)

	return nil
}

func SubscribeFromFile(client *reddit.Client, subPath string) (err error) {

	subredditNames, _ := scanLines(subPath)

	for _, v := range subredditNames {
		client.Subreddit.Subscribe(ctx, v)
		fmt.Printf("Subscribed to %v\n", v)
	}

	return nil
}

func Unsubscribe(client *reddit.Client, subredditName string) (err error) {

	client.Subreddit.Unsubscribe(ctx, subredditName)
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
