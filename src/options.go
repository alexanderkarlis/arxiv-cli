package arxivcli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Options struct {
	QueryString  string
	MaxResults   int
	SearchPrefix string
	OutputDir    string
}

func ParseOptions() *Options {
	var helpArg, helpSearchPrefix bool
	var queryTitle string
	var maxResults int
	var searchPrefix string
	var outputDir string

	flag.BoolVar(&helpArg, "-help", false, "--help")
	flag.BoolVar(&helpSearchPrefix, "prefixes", false, "-prefixes")
	flag.StringVar(&queryTitle, "t", "", "-t='electron'")
	flag.IntVar(&maxResults, "n", 10, "-n=3")
	flag.StringVar(&searchPrefix, "p", "all", "-p='au'")
	flag.StringVar(&outputDir, "o", "./", "-o='./output/directory/here/'")
	flag.Parse()

	if helpSearchPrefix {
		file, err := os.OpenFile("adds/prefixes.txt", os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}

		bytes, err := ioutil.ReadAll(file)
		fmt.Println("Search Prefix Keys ->")
		fmt.Println(string(bytes))
		os.Exit(1)
	}
	if helpArg {
		fmt.Println("arxiv-cli -- Usage: arxivcli [opts]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	return &Options{
		queryTitle,
		maxResults,
		searchPrefix,
		outputDir,
	}
}
