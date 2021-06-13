package client_wrapper

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type MockClient struct {
	Client *reddit.Client
}

func (r *MockClient) Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error) {
	s := []*reddit.Subreddit{
		{
			ID:     "1",
			FullID: "111",
			Name:   "My awesome subreddit",
		},
	}
	return s, nil, nil
}

func (r *MockClient) Subscribe(ctx context.Context, subreddit string) (*reddit.Response, error) {
	return nil, nil
}

func (r *MockClient) Unsubscribe(ctx context.Context, subreddit string) (*reddit.Response, error) {
	return nil, nil
}

func (r *MockClient) Upvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error) {
	u := []*reddit.Post{
		{
			ID:     "1",
			FullID: "111",
			URL:    "localhost:8080",
			Title:  "testPost",
		},
	}
	return u, nil, nil
}