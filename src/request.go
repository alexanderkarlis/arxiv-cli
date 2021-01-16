package arxivcli

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Response struct
type Response struct {
	Entries []Paper `xml:"entry"`
}

//Paper struct
type Paper struct {
	Title         string  `xml:"title"`
	ID            string  `xml:"id"`
	PublishedDate string  `xml:"published"`
	Summary       string  `xml:"summary"`
	Authors       Authors `xml:"author"`
}

// Author struct
type Author struct {
	Name        string `xml:"name"`
	Affiliation string `xml:"affiliation"`
}

// Authors struct
type Authors []Author

// String function returns the slice of Authors in string format
func (a *Authors) String() string {
	s := strings.Builder{}
	for _, author := range *a {
		s.WriteString(author.Name)
		s.WriteString(", ")
	}
	str := strings.TrimSuffix(s.String(), ", ")
	return str
}

// QueryRequest func
func QueryRequest(o *Options) *Response {
	urlBuilder := url.URL{
		Scheme: "https",
		Host:   "export.arxiv.org",
		Path:   "api/query",
	}
	q := urlBuilder.Query()
	q.Set("search_query", fmt.Sprintf("%s:%s", o.SearchPrefix, o.QueryString))
	q.Set("start", "0")
	q.Set("max_results", strconv.Itoa(o.MaxResults))
	urlBuilder.RawQuery = q.Encode()
	fmt.Println(urlBuilder.String())

	resp, err := http.Get(urlBuilder.String())
	b := []byte{}
	resp.Body.Read(b)
	fmt.Println(string(b))
	if err != nil {
		panic("bad request")
	}
	rs := Response{}
	// r := Response{}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rs)

	if err != nil {
		panic(err)
	}
	return &rs
}

func (p *Paper) Download(outDir string) (string, error) {
	urlBuilder := url.URL{
		Scheme: "http",
		Host:   "arxiv.org",
		Path:   "pdf",
	}

	split := strings.Split(p.ID, "/")
	paperID := split[len(split)-1]
	urlBuilder.Path = urlBuilder.Path + "/" + paperID + ".pdf"
	pdfPath := urlBuilder.String()

	// Get the data
	resp, err := http.Get(pdfPath)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	nsplit := strings.Split(p.Title, " ")
	strings.Join(nsplit, "")

	// Create the file
	fileName := strings.Join(nsplit, "") + ".pdf"
	fileName = strings.ReplaceAll(fileName, "'", "")
	fileName = strings.ReplaceAll(fileName, "\"", "")
	out, err := os.Create(outDir + fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return outDir + fileName, err
}
