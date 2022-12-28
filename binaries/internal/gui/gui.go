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
	"github.com/AHAOAHA/Annal/binaries/internal/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gui",
	Aliases: []string{"g"},
	Short:   "gui",
	Long:    `gui`,
	Run:     gui,
}

var (
	// topWindow   fyne.Window
	releaseFunc = []func(){releaseServeRTMP}
)

func gui(cmd *cobra.Command, args []string) {
	if err := initEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	defer release()

	a := app.NewWithID("annal.gui")
	a.SetIcon(icon)

	w := a.NewWindow(config.PROJECT)
	// topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	x := container.NewCenter(welcomeScreen(w))
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

	welcomeItem := fyne.NewMenuItem("welcome", func() {
		obj := welcomeScreen(w)
		w.SetContent(obj)
		w.Content().Refresh()
	})

	todosItem := fyne.NewMenuItem("todos", func() {
		obj := todoScreen(w)
		w.SetContent(obj)
		w.Content().Refresh()
	})
	servertmpItem := fyne.NewMenuItem("servertmp", func() {
		obj := servertmpScreen(w)
		w.SetContent(obj)
		w.Content().Refresh()
	})
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	helpMenu := fyne.NewMenu("Help",
		// fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About", func() {
			u, _ := url.Parse("https://ahaoaha.github.io/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", welcomeItem, todosItem, servertmpItem)

	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit"),
		helpMenu,
	)

	return main
}

func release() {
	for _, f := range releaseFunc {
		f()
	}
}
