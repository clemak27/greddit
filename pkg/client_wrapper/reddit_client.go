package client_wrapper

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type RedditClient struct {
	Client *reddit.Client
}

var ctx = context.Background()

func (r *RedditClient) Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error) {
	s, res, err := r.Client.Subreddit.Subscribed(ctx, opts)
	return s, res, err
}

func (r *RedditClient) Subscribe(ctx context.Context, subreddit string) (*reddit.Response, error) {
	res, err := r.Client.Subreddit.Subscribe(ctx, subreddit)
	return res, err
}

func (r *RedditClient) Unsubscribe(ctx context.Context, subreddit string) (*reddit.Response, error) {
	res, err := r.Client.Subreddit.Unsubscribe(ctx, subreddit)
	return res, err
}

func (r *RedditClient) Upvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error) {
	s, res, err := r.Client.User.Upvoted(ctx, opts)
	return s, res, err
}
