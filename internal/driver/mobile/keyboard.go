package mobile

import (
	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/driver/mobile"
	"github.com/Jaleel/fyne/internal/driver/mobile/app"
)

func showVirtualKeyboard(keyboard mobile.KeyboardType) {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		if driver.app == nil { // not yet running
			fyne.LogError("Cannot show keyboard before app is running", nil)
			return
		}

		driver.app.ShowVirtualKeyboard(app.KeyboardType(keyboard))
	}
}

func hideVirtualKeyboard() {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		if driver.app == nil { // not yet running
			return
		}

		driver.app.HideVirtualKeyboard()
	}
}
