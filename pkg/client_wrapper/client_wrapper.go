package client_wrapper

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type ClientFunctions interface {
	Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error)
	Upvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error)
}
