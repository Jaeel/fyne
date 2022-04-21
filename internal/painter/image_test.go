package painter_test

import (
	"testing"

	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/canvas"
	"github.com/Jaleel/fyne/internal/painter/software"
	"github.com/Jaleel/fyne/test"
)

func TestPaintImage_SVG(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		width     float32
		height    float32
		fillMode  canvas.ImageFill
		wantImage string
	}{
		"default": {
			width:  480,
			height: 240,
		},
		"stretchx": {
			width:    640,
			height:   240,
			fillMode: canvas.ImageFillStretch,
		},
		"stretchy": {
			width:    480,
			height:   480,
			fillMode: canvas.ImageFillStretch,
		},
		"containx": {
			width:    640,
			height:   240,
			fillMode: canvas.ImageFillContain,
		},
		"containy": {
			width:    480,
			height:   480,
			fillMode: canvas.ImageFillContain,
		},
	} {
		t.Run(name, func(t *testing.T) {
			img := canvas.NewImageFromFile("testdata/stroke.svg")
			c := test.NewCanvasWithPainter(software.NewPainter())
			c.SetContent(img)
			c.Resize(fyne.NewSize(tt.width, tt.height))
			img.Refresh()
			img.FillMode = tt.fillMode

			test.AssertImageMatches(t, "svg-stroke-"+name+".png", c.Capture())
		})
	}
}
