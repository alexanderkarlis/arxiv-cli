package main

import (
	"fmt"
	"os"

	arxivcli "github.com/alexanderkarlis/arxiv/src"
)

func main() {
	opts := arxivcli.ParseOptions()
	r := arxivcli.QueryRequest(opts)
	paper, err := arxivcli.ShowResults(r)
	if err != nil || paper.Title == "" {
		fmt.Println(err)
		os.Exit(1)
	}
	paper.Download()
}
