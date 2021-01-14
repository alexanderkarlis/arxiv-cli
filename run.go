package main

import (
	"flag"
	"fmt"
	"os"

	arxivcli "github.com/alexanderkarlis/arxiv/src"
)

func main() {
	var queryTitle string

	flag.StringVar(&queryTitle, "t", "black hole", "query arxiv for a specific string in the title.")
	flag.Parse()

	r := arxivcli.QueryRequest(queryTitle)
	paper, err := arxivcli.ShowResults(r)
	if err != nil || paper.Title == "" {
		fmt.Println(err)
		os.Exit(1)
	}
	paper.Download()
}
