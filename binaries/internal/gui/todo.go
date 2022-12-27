package gui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/AHAOAHA/Annal/binaries/internal/todo"
)

// todoMasterScreen loads a tab panel for collection widgets
func todoMasterScreen(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabelWithStyle(fmt.Sprintf("Total: %v", todo.Total()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Done: %v", todo.Done()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("Expired: %v", todo.Expired()), fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}))
	return container.NewCenter(content)
}

func makeListTab(_ fyne.Window) fyne.CanvasObject {
	tasks, err := storage.ListTodoTasks(context.Background(), storage.GetInstance())
	if err != nil {
		fyne.LogError("%v", err)
	}

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	hbox := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int {
			return len(tasks)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(tasks[id].GetTitle())
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(tasks[id].GetTitle())
		icon.SetResource(theme.DocumentIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	list.Select(125)
	list.SetItemHeight(5, 50)
	list.SetItemHeight(6, 50)

	return container.NewHSplit(list, container.NewCenter(hbox))
}
