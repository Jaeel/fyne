package mobile

import (
	"github.com/Jaeel/fyne/driver/mobile"
	"github.com/Jaeel/fyne/internal/driver/mobile/event/size"

	"github.com/Jaeel/fyne"
)

type device struct {
	safeTop, safeLeft, safeWidth, safeHeight int
}

//lint:file-ignore U1000 Var currentDPI is used in other files, but not here
var (
	currentOrientation size.Orientation
	currentDPI         float32
)

// Declare conformity with Device
var _ fyne.Device = (*device)(nil)

func (*device) Orientation() fyne.DeviceOrientation {
	switch currentOrientation {
	case size.OrientationLandscape:
		return fyne.OrientationHorizontalLeft
	default:
		return fyne.OrientationVertical
	}
}

func (*device) IsMobile() bool {
	return true
}

func (*device) HasKeyboard() bool {
	return false
}

func (*device) ShowVirtualKeyboard() {
	showVirtualKeyboard(mobile.DefaultKeyboard)
}

func (*device) ShowVirtualKeyboardType(keyboard mobile.KeyboardType) {
	showVirtualKeyboard(keyboard)
}

func (*device) HideVirtualKeyboard() {
	hideVirtualKeyboard()
}
