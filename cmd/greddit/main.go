package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"encoding/json"

	"github.com/vartanbeno/go-reddit/v2/reddit"
	"gitlab.com/clemak27/greddit/pkg/authentication"
)

var ctx = context.Background()

func main() {
	credentials, _ := getConfig()
	authentication.GetClient(credentials)
}

func getConfig() (credentials reddit.Credentials, err error) {

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &credentials)

	return credentials, nil
}
