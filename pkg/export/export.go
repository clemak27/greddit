package export

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/vartanbeno/go-reddit/v2/reddit"
	"gitlab.com/clemak27/greddit/pkg/upvoted"
)

var ctx = context.Background()

func ExportUpvoted(client *reddit.Client, format string) (err error) {

	l := upvoted.GetUpvoted(client)
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
		generateMdFile(res)
	case "html":
		generateHTMLFile(res)
	case "txt":
		generateTxtFile(res)
	default:
		fmt.Printf("Unknown output format %s! Supported formats are: md, html", format)
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

func generateMdFile(res map[string][]reddit.Post) {
	fn := "./pkg/export/md.tmpl"
	ofn := "./export-upvoted.md"

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

func generateHTMLFile(res map[string][]reddit.Post) {
	fn := "./pkg/export/html.tmpl"
	ofn := "./export-upvoted.html"

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

func generateTxtFile(res map[string][]reddit.Post) {
	fn := "./pkg/export/txt.tmpl"
	ofn := "./export-upvoted.txt"

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
