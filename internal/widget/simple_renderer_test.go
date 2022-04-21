package widget_test

import (
	"image/color"
	"testing"

	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/canvas"
	"github.com/Jaleel/fyne/internal/widget"
	"github.com/Jaleel/fyne/test"
)

func TestNewSimpleRenderer(t *testing.T) {
	r := canvas.NewRectangle(color.Transparent)
	o := &simpleWidget{obj: r}
	o.ExtendBaseWidget(o)
	w := test.NewWindow(o)
	w.Resize(fyne.NewSize(100, 100))

	test.AssertRendersToMarkup(t, "simple_renderer.xml", w.Canvas())
}

type simpleWidget struct {
	widget.Base
	obj fyne.CanvasObject
}

func (s *simpleWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.obj)
}
