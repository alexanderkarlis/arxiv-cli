package arxivcli

import (
	"flag"
)

type Options struct {
	Title      string
	MaxResults int
}

func ParseOptions() *Options {
	var queryTitle string
	var maxResults int
	flag.StringVar(&queryTitle, "t", "black hole", "query arxiv for a specific string in the title.")
	flag.IntVar(&maxResults, "n", 10, "-n=3")
	flag.Parse()

	return &Options{
		queryTitle,
		maxResults,
	}

}
