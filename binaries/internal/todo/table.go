package todo

import (
	"os"
	"time"

	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/jedib0t/go-pretty/table"
)

type TableOptions struct {
	Columns []int
	SortBy  int
	Style   table.Style
}

type Column struct {
	ID        string
	Name      string
	SortIndex int
	Width     int
}

var columns = []Column{
	{ID: "id", Name: "ID", SortIndex: 1},
	{ID: "uuid", Name: "UUID", SortIndex: 12, Width: 7},
	{ID: "title", Name: "Title", SortIndex: 13, Width: 7},
	{ID: "description", Name: "Description", SortIndex: 14, Width: 7},
	{ID: "plan", Name: "Plan", SortIndex: 15, Width: 6},
	{ID: "status", Name: "Status", SortIndex: 16, Width: 7},
	{ID: "created_at", Name: "CreatedAt", SortIndex: 17, Width: 7},
	{ID: "updated_at", Name: "UpdatedAt", SortIndex: 18, Width: 7},
}

// printTable prints an individual table of mounts.
func printTable(title string, m []*storage.TodoTask, opts TableOptions) {
	tab := table.NewWriter()
	tab.SetOutputMirror(os.Stdout)

	headers := table.Row{}
	for _, v := range columns {
		headers = append(headers, v.Name)
	}
	tab.AppendHeader(headers)

	for _, v := range m {
		tab.AppendRow([]interface{}{
			v.ID,
			v.UUID,
			v.Title,
			v.Description,
			v.Plan,
			v.Status,
			time.Unix(v.CreatedAt, 0).String(),
			time.Unix(v.UpdatedAt, 0).String(),
		})
	}

	tab.Render()
}
