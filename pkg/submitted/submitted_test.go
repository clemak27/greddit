package submitted

import (
	"reflect"
	"testing"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestGetSubmitted(t *testing.T) {
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
			name: "GetSubmitted is successfull",
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
			gotL, err := GetSubmitted(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubmitted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetSubmitted() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}

func TestGetSubmittedComments(t *testing.T) {
	type args struct {
		rc client_wrapper.ClientWrapper
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Comment
		wantErr bool
	}{
		{
			name: "GetSubmittedComments is successfull",
			args: args{
				rc: &client_wrapper.MockClient{},
			},
			wantL: []*reddit.Comment{
				{
					ID:        "62",
					FullID:    "6254",
					PostTitle: "Some post about stuff",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := GetSubmittedComments(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubmittedComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetSubmittedComments() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
