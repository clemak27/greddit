package authentication

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func GetClient(credentials reddit.Credentials) (client *reddit.Client, err error) {

	client, err = reddit.NewClient(credentials)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return
	}

	user, _, err := client.Account.Info(ctx)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return
	}

	fmt.Println("Authenticated as", user.Name)

	return client, nil
}
