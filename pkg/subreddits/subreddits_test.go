package subreddits

import (
	"reflect"
	"testing"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestPrintSubcriptions(t *testing.T) {
	type args struct {
		client *reddit.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintSubcriptions(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("PrintSubcriptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSubscriptions(t *testing.T) {
	type args struct {
		client *reddit.Client
	}
	tests := []struct {
		name    string
		args    args
		wantL   []*reddit.Subreddit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := GetSubscriptions(tt.args.client)
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
		client *reddit.Client
		name   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Subscribe(tt.args.client, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscribeFromFile(t *testing.T) {
	type args struct {
		client  *reddit.Client
		subPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SubscribeFromFile(tt.args.client, tt.args.subPath); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	type args struct {
		client        *reddit.Client
		subredditName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unsubscribe(tt.args.client, tt.args.subredditName); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scanLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("scanLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retrieveMore(t *testing.T) {
	type args struct {
		subs   []*reddit.Subreddit
		client *reddit.Client
	}
	tests := []struct {
		name string
		args args
		want []*reddit.Subreddit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retrieveMore(tt.args.subs, tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveMore() = %v, want %v", got, tt.want)
			}
		})
	}
}
