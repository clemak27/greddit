package export

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/clemak27/greddit/pkg/client_wrapper"
	"github.com/clemak27/greddit/pkg/upvoted"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

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

	switch format {
	case "md":
		fn := "./pkg/export/md.tmpl"
		ofn := "./export-upvoted.md"
		generateOutputFile(res, fn, ofn)
	case "html":
		fn := "./pkg/export/html.tmpl"
		ofn := "./export-upvoted.html"
		generateOutputFile(res, fn, ofn)
	case "txt":
		fn := "./pkg/export/txt.tmpl"
		ofn := "./export-upvoted.txt"
		generateOutputFile(res, fn, ofn)
	default:
		fmt.Printf("Unknown output format %s! Supported formats are: md, html, txt", format)
	}

	return nil
}

func printResult(res map[string][]reddit.Post) {
	for k := range res {
		fmt.Println(k)
		for _, v := range res[k] {
			fmt.Printf("  %s\n", v.Title)
		}
	}
}

func generateOutputFile(res map[string][]reddit.Post, fn string, ofn string) {

	tpl, err := template.ParseFiles(fn)
	if err != nil {
		fmt.Println("Failed to parse template")
	}

	f, err := os.Create(ofn)
	if err != nil {
		fmt.Println("Failed to open output file!")
	}

	err = tpl.Execute(f, res)
	if err != nil {
		fmt.Println("Failed to write output file!")
	}

	fmt.Printf("Wrote output to %s!", ofn)
}
