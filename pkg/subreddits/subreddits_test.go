package subreddits

import (
	"context"
	"reflect"
	"testing"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type mockClient struct {
	Client *reddit.Client
}

func (r *mockClient) Subs(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error) {
	s := []*reddit.Subreddit{
		{
			ID:     "1",
			FullID: "111",
			Name:   "My awesome subreddit",
		},
	}
	return s, nil, nil
}

func TestGetSubscriptionsNew(t *testing.T) {
	type args struct {
		rc client_wrapper.RedditClient
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Subreddit
		wantErr bool
	}{
		{
			name: "succ",
			args: args{
				rc: &mockClient{},
			},
			wantL: []*reddit.Subreddit{
				{
					ID:     "1",
					FullID: "111",
					Name:   "My awesome subreddit",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := GetSubscriptions(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscriptionsNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetSubscriptionsNew() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
