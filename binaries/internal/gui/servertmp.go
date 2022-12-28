package gui

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/AHAOAHA/Annal/binaries/internal/rtmp"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
)

func servertmpScreen(w fyne.Window) fyne.CanvasObject {
	addrEntry := widget.NewEntry()
	addrEntry.SetPlaceHolder("input server address:")
	addrEntry.OnChanged = func(content string) {
		fmt.Println("address:", addrEntry.Text, "entered")
	}
	addrEntry.Resize(fyne.NewSize(200, 37))

	startBtn := widget.NewButton("Start", func() {
		mu.Lock()
		defer mu.Unlock()
		if ctx != nil {
			fmt.Fprintf(os.Stderr, "%s", "already in serve")
			return
		}

		ctx, cancel = context.WithCancel(context.Background())
		go func() {
			rtmp.ServeRTMP(ctx, addrEntry.Text)
		}()
	})
	stopBtn := widget.NewButton("Stop", func() {
		mu.Lock()
		defer mu.Unlock()
		if ctx == nil {
			fmt.Fprintf(os.Stderr, "%s\n", "not in serve")
			return
		}

		cancel()
		select {
		case <-ctx.Done():
			ctx = nil
			cancel = nil
		case <-time.After(time.Second * 3):
			fmt.Fprintf(os.Stderr, "%s\n", "timeout")
		}

	})
	btns := container.NewHBox(startBtn, stopBtn)
	btns.Move(fyne.NewPos(45, 0))

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to Annal ServeRTMP", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewWithoutLayout(addrEntry),
		container.NewWithoutLayout(btns),
		widget.NewLabel("")))
}
