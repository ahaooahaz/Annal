package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/AHAOAHA/Annal/binaries/internal/todo"
)

// todoMasterScreen loads a tab panel for collection widgets
func todoScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle(fmt.Sprintf("Total: %v", todo.Total()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Done: %v", todo.Done()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Expired: %v", todo.Expired()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}))
}
