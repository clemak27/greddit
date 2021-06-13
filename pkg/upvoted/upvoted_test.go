package upvoted

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

func (r *mockClient) Subscribed(ctx context.Context, opts *reddit.ListSubredditOptions) ([]*reddit.Subreddit, *reddit.Response, error) {
	return nil, nil, nil
}

func (r *mockClient) Upvoted(ctx context.Context, opts *reddit.ListUserOverviewOptions) ([]*reddit.Post, *reddit.Response, error) {
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

func TestGetUpvoted(t *testing.T) {
	type args struct {
		rc client_wrapper.ClientFunctions
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Post
		wantErr bool
	}{
		{
			name: "GetUpvoted is successfull",
			args: args{
				rc: &mockClient{},
			},
			wantL: []*reddit.Post{
				{
					ID:     "1",
					FullID: "111",
					URL:    "localhost:8080",
					Title:  "testPost",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := GetUpvoted(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpvoted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetUpvoted() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
