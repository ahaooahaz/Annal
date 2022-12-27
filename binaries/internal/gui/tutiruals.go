package gui

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	_tutorials = map[string]Tutorial{
		"welcome": {"Welcome", "", welcomeScreen, true},
		"todo": {"Todo",
			"todo tasks",
			todoMasterScreen, // todo.TodoScreen,
			true,
		},
		"list": {"List",
			"list all tasks.",
			makeListTab,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	_tutorialIndex = map[string][]string{
		"": {"welcome"},
	}
)
