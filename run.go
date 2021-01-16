package main

import (
	"fmt"
	"os"
	"os/exec"

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
	filename, err := paper.Download(opts.OutputDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = exec.Command("xdg-open", filename).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
