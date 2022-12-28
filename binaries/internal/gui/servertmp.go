package gui

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	cmd *exec.Cmd
	mu  sync.Mutex
)

func servertmpScreen(w fyne.Window) fyne.CanvasObject {
	addrEntry := widget.NewEntry()
	addrEntry.SetPlaceHolder("server address default (:1935)")
	addrEntry.OnChanged = func(content string) {
		fmt.Println("address:", addrEntry.Text, "entered")
	}
	addrEntry.Resize(fyne.NewSize(210, 36))

	startBtn := widget.NewButton("Start", func() {
		mu.Lock()
		defer mu.Unlock()
		if cmd != nil {
			fmt.Fprintf(os.Stderr, "%s", "already in serve")
			return
		}

		cmd = exec.Command("annal", "servertmp", addrEntry.Text)
		cmd.Env = os.Environ()
		err := cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			return
		}
	})

	stopBtn := widget.NewButton("Stop", func() {
		mu.Lock()
		defer mu.Unlock()
		if cmd == nil {
			fmt.Fprintf(os.Stderr, "%s\n", "not in serve")
			return
		}

		cmd.Process.Kill()
		cmd.Wait()
		cmd = nil
	})
	btns := container.NewHBox(startBtn, stopBtn)
	btns.Move(fyne.NewPos(50, 0))

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to Annal ServeRTMP", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewWithoutLayout(addrEntry),
		widget.NewLabel(""),
		container.NewWithoutLayout(btns),
		widget.NewLabel("")))
}

func releaseServeRTMP() {
	if cmd != nil {
		cmd.Process.Kill()
		cmd.Wait()
	}
}
