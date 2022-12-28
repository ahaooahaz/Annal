package gui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/AHAOAHA/Annal/binaries/internal/todo"
)

// todoMasterScreen loads a tab panel for collection widgets
func todoScreen(_ fyne.Window) fyne.CanvasObject {
	s := todo.Statistic(context.TODO())
	return container.NewVBox(
		widget.NewLabelWithStyle(fmt.Sprintf("Total: %v", s.Total), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Done: %v", s.Done), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Expired: %v", s.Expired), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}))
}
