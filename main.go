package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	InitWebRTC()
	a := app.New()
	w := a.NewWindow("Hello")
	RV := NewRemoteView()
	go captureScreen(RV)
	w.SetContent(RV)
	w.ShowAndRun()
}
