package client_wrapper

import (
	"context"
	"testing"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestSubs(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts reddit.ListSubredditOptions
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Subs(tt.args.ctx, tt.args.opts)
		})
	}
}
