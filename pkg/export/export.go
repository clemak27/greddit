package export

import (
	"context"
	"errors"
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

type commentContent struct {
	Title string
	Items []*reddit.Comment
}

func Posts(rc client_wrapper.ClientWrapper, format string, tp string) (err error) {

	var l []*reddit.Post

	switch tp {
	case "upvoted":
		var err error
		l, err = upvoted.GetUpvoted(rc)
		if err != nil {
			fmt.Println("Failed to get upvoted posts!")
			return err
		}
	case "saved":
		var err error
		l, err = saved.GetSaved(rc)
		if err != nil {
			fmt.Println("Failed to get saved posts!")
			return err
		}
	case "downvoted":
		var err error
		l, err = downvoted.GetDownvoted(rc)
		if err != nil {
			fmt.Println("Failed to get downvoted posts!")
			return err
		}
	case "submitted":
		var err error
		l, err = submitted.GetSubmitted(rc)
		if err != nil {
			fmt.Println("Failed to get submitted posts!")
			return err
		}
	default:
		fmt.Printf("Unknown type %s!", tp)
		return err
	}

	cont := content{
		Title: fmt.Sprintf("%v reddit posts", tp),
		Items: listOfPosts(l),
	}

	switch format {
	case "md", "html", "txt":
		fn := fmt.Sprintf("./pkg/export/%v.tmpl", format)
		ofn := fmt.Sprintf("./export-%v.%v", tp, format)
		err := generateOutputFile(cont, fn, ofn)
		if err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf("Unknown output format %s! Supported formats are: md, html, txt", format)
		return errors.New(msg)
	}

	return nil
}

func Comments(rc client_wrapper.ClientWrapper, format string, tp string) (err error) {

	var l []*reddit.Comment

	switch tp {
	case "saved-comments":
		var err error
		l, err = saved.GetSavedComments(rc)
		if err != nil {
			fmt.Println("Failed to get saved comments!")
			return err
		}
	case "submitted-comments":
		var err error
		l, err = submitted.GetSubmittedComments(rc)
		if err != nil {
			fmt.Println("Failed to get saved comments!")
			return err
		}
	default:
		fmt.Printf("Unknown type %s!", tp)
	}

	cont := commentContent{
		Title: fmt.Sprintf("saved reddit comments"),
		Items: l,
	}

	switch format {
	case "md":
		fn := fmt.Sprintf("./pkg/export/comment-%v.tmpl", format)
		ofn := fmt.Sprintf("./export-%v.%v", tp, format)
		err := generateCommentOutputFile(cont, fn, ofn)
		if err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf("Unknown output format %s! Currently only md is supported", format)
		return errors.New(msg)
	}

	return nil
}

func generateOutputFile(cont content, fn string, ofn string) (err error) {

	tpl, err := template.ParseFiles(fn)
	if err != nil {
		fmt.Println("Failed to parse template")
		return err
	}

	f, err := os.Create(ofn)
	if err != nil {
		fmt.Println("Failed to open output file!")
		return err
	}

	err = tpl.Execute(f, cont)
	if err != nil {
		fmt.Println("Failed to write output file!")
		return err
	}

	fmt.Printf("Wrote output to %s!", ofn)
	return nil
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

func generateCommentOutputFile(cont commentContent, fn string, ofn string) (err error) {

	tpl, err := template.ParseFiles(fn)
	if err != nil {
		fmt.Println("Failed to parse template")
		return err
	}

	f, err := os.Create(ofn)
	if err != nil {
		fmt.Println("Failed to open output file!")
		return err
	}

	err = tpl.Execute(f, cont)
	if err != nil {
		fmt.Println("Failed to write output file!")
		return err
	}

	fmt.Printf("Wrote output to %s!", ofn)
	return nil
}
