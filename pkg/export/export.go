package export

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/clemak27/greddit/pkg/saved"
	"github.com/clemak27/greddit/pkg/upvoted"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

type content struct {
	Title string
	Items map[string][]reddit.Post
}

func ExportUpvoted(rc client_wrapper.ClientWrapper, format string) (err error) {

	l, err := upvoted.GetUpvoted(rc)
	var res = make(map[string][]reddit.Post, 0)

	for _, p := range l {
		sr, exists := res[p.SubredditName]
		if exists {
			res[p.SubredditName] = append(sr, *p)
		} else {
			res[p.SubredditName] = []reddit.Post{*p}
		}
	}

	cont := content{
		Title: "upvoted Reddit posts",
		Items: res,
	}

	switch format {
	case "md":
		fn := "./pkg/export/md.tmpl"
		ofn := "./export-upvoted.md"
		generateOutputFile(cont, fn, ofn)
	case "html":
		fn := "./pkg/export/html.tmpl"
		ofn := "./export-upvoted.html"
		generateOutputFile(cont, fn, ofn)
	case "txt":
		fn := "./pkg/export/txt.tmpl"
		ofn := "./export-upvoted.txt"
		generateOutputFile(cont, fn, ofn)
	default:
		fmt.Printf("Unknown output format %s! Supported formats are: md, html, txt", format)
	}

	return nil
}

func ExportSaved(rc client_wrapper.ClientWrapper, format string) (err error) {

	l, err := saved.GetSaved(rc)
	var res = make(map[string][]reddit.Post, 0)

	for _, p := range l {
		sr, exists := res[p.SubredditName]
		if exists {
			res[p.SubredditName] = append(sr, *p)
		} else {
			res[p.SubredditName] = []reddit.Post{*p}
		}
	}

	cont := content{
		Title: "saved Reddit posts",
		Items: res,
	}

	switch format {
	case "md":
		fn := "./pkg/export/md.tmpl"
		ofn := "./export-saved.md"
		generateOutputFile(cont, fn, ofn)
	case "html":
		fn := "./pkg/export/html.tmpl"
		ofn := "./export-saved.html"
		generateOutputFile(cont, fn, ofn)
	case "txt":
		fn := "./pkg/export/txt.tmpl"
		ofn := "./export-saved.txt"
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
