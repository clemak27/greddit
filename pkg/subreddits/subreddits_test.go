package subreddits

import (
	"reflect"
	"testing"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestGetSubscriptions(t *testing.T) {
	type args struct {
		rc client_wrapper.ClientWrapper
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Subreddit
		wantErr bool
	}{
		{
			name: "GetSubscriptions is successfull",
			args: args{
				rc: &client_wrapper.MockClient{},
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
				t.Errorf("GetSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("GetSubscriptions() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	type args struct {
		rc   client_wrapper.ClientWrapper
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Subscribe succesfully",
			args: args{
				rc:   &client_wrapper.MockClient{},
				name: "testSubreddit",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Subscribe(tt.args.rc, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscribeFromFile(t *testing.T) {
	type args struct {
		rc      client_wrapper.ClientWrapper
		subPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Subscribe succesfully",
			args: args{
				rc:      &client_wrapper.MockClient{},
				subPath: "./subs.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SubscribeFromFile(tt.args.rc, tt.args.subPath); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	type args struct {
		rc            client_wrapper.ClientWrapper
		subredditName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Unsubscribe succesfully",
			args: args{
				rc:            &client_wrapper.MockClient{},
				subredditName: "testSubreddit",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unsubscribe(tt.args.rc, tt.args.subredditName); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
