package gui

import (
	"fmt"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/AHAOAHA/Annal/binaries/internal/config"
	"github.com/spf13/cobra"
)

const preferenceCurrentTutorial = "currentTutorial"

var Cmd = &cobra.Command{
	Use:     "gui",
	Aliases: []string{"g"},
	Short:   "gui",
	Long:    `gui`,
	Run:     gui,
}

var topWindow fyne.Window

func gui(cmd *cobra.Command, args []string) {
	if err := initEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	a := app.NewWithID("annal.gui")
	a.SetIcon(icon)

	w := a.NewWindow(config.PROJECT)
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	content := container.NewMax()
	intro := widget.NewLabel("intro")
	intro.Wrapping = fyne.TextWrapWord
	setTutorial := func(t Tutorial) {
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(intro), nil, nil, nil, content)

	x := container.NewCenter(makeNav(setTutorial, true), tutorial)
	w.SetContent(x)

	w.Resize(fyne.NewSize(640, 460))

	w.ShowAndRun()

}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	todosItem := fyne.NewMenuItem("todos", func() {
		// TODO: action todos
	})
	servertmpItem := fyne.NewMenuItem("servertmp", func() {
		// TODO: action servertmp
	})
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		// fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About", func() {
			u, _ := url.Parse("https://ahaoaha.github.io/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", todosItem, servertmpItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)

	return main
}

func unsupportedTutorial(t Tutorial) bool {
	return !t.SupportWeb && fyne.CurrentDevice().IsBrowser()
}

func makeNav(setTutorial func(tutorial Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return _tutorialIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := _tutorialIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := _tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			if unsupportedTutorial(t) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			} else {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			}
		},
		OnSelected: func(uid string) {
			if t, ok := _tutorials[uid]; ok {
				if unsupportedTutorial(t) {
					return
				}
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	return container.NewBorder(nil, nil, nil, nil)
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
