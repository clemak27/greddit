package client_wrapper

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type ClientFunctions interface {
	Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error)
}

type RedditClient struct {
	Client *reddit.Client
}

var ctx = context.Background()

func (r *RedditClient) Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error) {
	s, res, err := r.Client.Subreddit.Subscribed(ctx, opts)
	return s, res, err
}
