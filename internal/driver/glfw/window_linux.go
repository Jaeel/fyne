package glfw

import "github.com/Jaeel/fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
	w.canvas.Refresh(w.canvas.content)
}
