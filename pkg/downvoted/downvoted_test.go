package downvoted

import (
	"reflect"
	"testing"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestGetDownvoted(t *testing.T) {
	type args struct {
		rc client_wrapper.ClientWrapper
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Post
		wantErr bool
	}{
		{
			name: "GetDownvoted is successfull",
			args: args{
				rc: &client_wrapper.MockClient{},
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
			gotL, err := GetDownvoted(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDownvoted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetDownvoted() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
