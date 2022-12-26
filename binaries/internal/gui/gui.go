package gui

import (
	"fmt"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gui",
	Aliases: []string{"g"},
	Short:   "gui task",
	Long:    `gui task`,
	Run:     gui,
}

func gui(cmd *cobra.Command, args []string) {
	if err := initEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	app := app.New()
	w := app.NewWindow("Hello")
	w.SetIcon(icon)
	w.SetFullScreen(true)
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("todo", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()

}
