package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

type RemoteView struct {
	widget.BaseWidget
	screen *canvas.Image
}

func (RV *RemoteView) FocusGained() {
	fmt.Println("focus gained")
}

func (RV *RemoteView) FocusLost() {
	fmt.Println("focus lost")
}

func (RV *RemoteView) TypedRune(r rune) {
	fmt.Printf("typed rune %c\n", r)
}

func (RV *RemoteView) TypedKey(e *fyne.KeyEvent) {
	fmt.Printf("typed key %v\n", e.Name)
}
func (RV *RemoteView) Tapped(*fyne.PointEvent) {
	fyne.CurrentApp().Driver().CanvasForObject(RV).Focus(RV)
}

func (RV *RemoteView) MouseMoved(e *desktop.MouseEvent) {
	fmt.Printf("Mouse moved: X=%.4f Y=%.4f\n", e.Position.X, e.Position.Y)
}
func (RV *RemoteView) MouseIn(e *desktop.MouseEvent) {
	fmt.Println("mouse in")
}
func (RV *RemoteView) MouseOut() {
	fmt.Println("mouse out")
}

func (RV *RemoteView) KeyDown(e *fyne.KeyEvent) {
	fmt.Printf("keydown %v\n", e.Name)
}

func (RV *RemoteView) KeyUp(e *fyne.KeyEvent) {
	fmt.Printf("keyup %v\n", e.Name)
}
func (RV *RemoteView) CreateRenderer() fyne.WidgetRenderer {
	RV.screen = canvas.NewImageFromImage(nil)
	RV.screen.FillMode = canvas.ImageFillContain
	RV.screen.SetMinSize(fyne.NewSize(800, 600))

	rect := canvas.NewRectangle(color.Black)
	return widget.NewSimpleRenderer(
		container.NewStack(rect, RV.screen),
	)
}
func (RV *RemoteView) SetFrame(img image.Image) {
	if RV.screen == nil {
		return // renderer not ready yet
	}

	RV.screen.Image = img
	RV.screen.Refresh()
}

func NewRemoteView() *RemoteView {
	RV := &RemoteView{}
	RV.ExtendBaseWidget(RV)
	return RV
}

func captureScreen(RV *RemoteView) {
	for {
		bounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fyne.Do(func() { RV.SetFrame(img) })
		time.Sleep(33 * time.Millisecond) // ~30 FPS

	}
}
