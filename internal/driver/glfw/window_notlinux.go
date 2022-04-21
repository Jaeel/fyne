// +build !linux

package glfw

import "github.com/Jaleel/fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	d, ok := fyne.CurrentApp().Driver().(*gLDriver)
	if !ok { // don't wait to redraw in this way if we are running on test
		w.canvas.Resize(canvasSize)
		return
	}

	runOnDraw(w, func() {
		w.canvas.Resize(canvasSize)
		d.repaintWindow(w)
	})
}
