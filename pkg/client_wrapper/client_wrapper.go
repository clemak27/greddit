package client_wrapper

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type ClientWrapper interface {
	Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error)
	Subscribe(ctx context.Context, subreddit string) (*reddit.Response, error)
	Unsubscribe(ctx context.Context, subreddit string) (*reddit.Response, error)
	Upvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error)
	Saved(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, []*reddit.Comment, *reddit.Response, error)
	Downvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error)
	Submitted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error)
	SubmittedComments(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Comment, *reddit.Response, error)
}
