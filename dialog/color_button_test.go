package dialog

import (
	"image/color"
	"testing"

	"github.com/Jaeel/fyne"
	"github.com/Jaeel/fyne/layout"
	"github.com/Jaeel/fyne/test"
	"github.com/Jaeel/fyne/theme"
)

func Test_colorButton_Layout(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		color   color.Color
		hovered bool
	}{
		"primary": {
			color: theme.PrimaryColor(),
		},
		"primary_hovered": {
			color:   theme.PrimaryColor(),
			hovered: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			color := newColorButton(tt.color, nil)

			if tt.hovered {
				color.MouseIn(nil)
			}

			window := test.NewWindow(fyne.NewContainerWithLayout(layout.NewCenterLayout(), color))
			window.Resize(color.MinSize().Max(fyne.NewSize(50, 50)))

			test.AssertImageMatches(t, "color/button_layout_"+name+".png", window.Canvas().Capture())

			window.Close()
		})
	}
}
