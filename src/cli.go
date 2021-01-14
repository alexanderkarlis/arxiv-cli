package arxivcli

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ShowPaperDesc(s Paper, x, y int) (*widgets.Paragraph, error) {
	p := widgets.NewParagraph()
	parsedDate, err := time.Parse(time.RFC3339, s.PublishedDate)
	if err != nil {
		return nil, err
	}
	strDate := parsedDate.Format("January 2, 2006")
	p.Text = fmt.Sprintf("%s\n%s\n\n%s\n\n%s", s.Authors.String(), strDate, s.Summary, s.ID)
	p.SetRect(x, 0, 2*x, y)
	return p, nil
}

func ShowResults(results *Response) (Paper, error) {
	if err := ui.Init(); err != nil {
		return Paper{}, err
	}
	defer ui.Close()
	l := widgets.NewList()
	l.Title = "Arxiv papers"
	rows := []string{}
	for i, r := range results.Entries {
		rows = append(rows, fmt.Sprintf("[%d] \"%s\"", i, r.Title))
	}
	if len(rows) == 0 {
		rows = append(rows, "No results found.")
	}
	l.Rows = rows
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	termWidth, termHeight := ui.TerminalDimensions()
	l.SetRect(0, 0, termWidth/2, termHeight)
	l.SelectedRowStyle = ui.Style{Bg: ui.ColorCyan}

	activeKey := 0
	previousKey := ""

	p, err := ShowPaperDesc(results.Entries[activeKey], termWidth/2, termHeight)
	if err != nil {
		return Paper{}, err
	}
	ui.Render(l, p)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return Paper{}, nil
		case "<Enter>":
			return results.Entries[l.SelectedRow], nil
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		p, _ := ShowPaperDesc(results.Entries[l.SelectedRow], termWidth/2, termHeight)
		ui.Render(l, p)
	}
}
