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

func (r *RedditClient) Saved(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, []*reddit.Comment, *reddit.Response, error) {
	s, c, res, err := r.Client.User.Saved(ctx, opts)
	return s, c, res, err
}

func (r *RedditClient) Downvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error) {
	s, res, err := r.Client.User.Downvoted(ctx, opts)
	return s, res, err
}

func (r *RedditClient) Submitted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error) {
	s, res, err := r.Client.User.Posts(ctx, opts)
	return s, res, err
}

func (r *RedditClient) SubmittedComments(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Comment, *reddit.Response, error) {
	s, res, err := r.Client.User.Comments(ctx, opts)
	return s, res, err
}
