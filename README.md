# Arxiv cli
Download papers from cli with search feature!

## How to use
```sh
> go run run.go -t="electron" -n=2

```
A term-ui will open up where you can select the papers that are return from the query. Select
`<Enter>` to download the paper in the local directory as a PDF. Press `<q>` to quit the term-ui
without downloading a paper.

## Options
| Option       | cli-arg | example       |
| :------------- | :----------: | -----------: |
| title        | -t | -t="electron" |
| max-results  | -n | -n=10         |

__More to come!!!!__