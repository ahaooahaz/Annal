package todo

import (
	"os"
	"time"

	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

type TableOptions struct {
	Columns []Column
	SortBy  int
	Style   table.Style
}

type Column struct {
	ID        string
	Name      string
	SortIndex int
	Width     int
	colors    []text.Color
}

var columns = []Column{
	{ID: "id", Name: "ID", colors: []text.Color{text.BgYellow}},
	// {ID: "uuid", Name: "UUID", colors: []text.Color{text.BgCyan}},
	{ID: "index", Name: "Index", colors: []text.Color{text.BgCyan}},
	{ID: "title", Name: "Title"},
	{ID: "description", Name: "Description"},
	{ID: "plan", Name: "Plan", colors: []text.Color{text.BgHiRed}},
	{ID: "status", Name: "Status", colors: []text.Color{text.BgHiGreen}},
	{ID: "created_at", Name: "CreatedAt", colors: []text.Color{text.BgYellow}},
	{ID: "updated_at", Name: "UpdatedAt", colors: []text.Color{text.BgYellow}},
}

// printTable prints an individual table of mounts.
func printTable(title string, m []*pb.TodoTask, opts TableOptions) {
	tab := table.NewWriter()
	tab.SetOutputMirror(os.Stdout)
	tab.SetStyle(opts.Style)

	colors := []text.Colors{}
	headers := table.Row{}
	for _, v := range opts.Columns {
		headers = append(headers, v.Name)
		colors = append(colors, v.colors)
	}
	tab.AppendHeader(headers)
	tab.SetColors(colors)

	for _, v := range m {
		tab.AppendRow([]interface{}{
			v.ID,
			// v.UUID,
			v.Index,
			v.Title,
			v.Description,
			time.Unix(v.Plan, 0).Format(_TimeFormatString),
			v.Status,
			time.Unix(v.CreatedAt, 0).Format(_TimeFormatString),
			time.Unix(v.UpdatedAt, 0).Format(_TimeFormatString),
		})
	}

	tab.Render()
}
