package export

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/clemak27/greddit/pkg/downvoted"
	"github.com/clemak27/greddit/pkg/saved"
	"github.com/clemak27/greddit/pkg/submitted"
	"github.com/clemak27/greddit/pkg/upvoted"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

type content struct {
	Title string
	Items map[string][]reddit.Post
}

func Posts(rc client_wrapper.ClientWrapper, format string, tp string) (err error) {

	var l []*reddit.Post

	switch tp {
	case "upvoted":
		var err error
		l, err = upvoted.GetUpvoted(rc)
		if err != nil {
			fmt.Println("Failed to get upvoted posts!")
		}
	case "saved":
		var err error
		l, err = saved.GetSaved(rc)
		if err != nil {
			fmt.Println("Failed to get saved posts!")
		}
	case "downvoted":
		var err error
		l, err = downvoted.GetDownvoted(rc)
		if err != nil {
			fmt.Println("Failed to get downvoted posts!")
		}
	case "submitted":
		var err error
		l, err = submitted.GetSubmitted(rc)
		if err != nil {
			fmt.Println("Failed to get submitted posts!")
		}
	default:
		fmt.Printf("Unknown type %s!", tp)
	}

	cont := content{
		Title: fmt.Sprintf("%v reddit posts", tp),
		Items: listOfPosts(l),
	}

	switch format {
	case "md", "html", "txt":
		fn := fmt.Sprintf("./pkg/export/%v.tmpl", format)
		ofn := fmt.Sprintf("./export-%v.%v", tp, format)
		generateOutputFile(cont, fn, ofn)
	default:
		fmt.Printf("Unknown output format %s! Supported formats are: md, html, txt", format)
	}

	return nil
}

func generateOutputFile(cont content, fn string, ofn string) {

	tpl, err := template.ParseFiles(fn)
	if err != nil {
		fmt.Println("Failed to parse template")
	}

	f, err := os.Create(ofn)
	if err != nil {
		fmt.Println("Failed to open output file!")
	}

	err = tpl.Execute(f, cont)
	if err != nil {
		fmt.Println("Failed to write output file!")
	}

	fmt.Printf("Wrote output to %s!", ofn)
}

func listOfPosts(l []*reddit.Post) map[string][]reddit.Post {
	var res = make(map[string][]reddit.Post, 0)

	for _, p := range l {
		sr, exists := res[p.SubredditName]
		if exists {
			res[p.SubredditName] = append(sr, *p)
		} else {
			res[p.SubredditName] = []reddit.Post{*p}
		}
	}

	return res
}
