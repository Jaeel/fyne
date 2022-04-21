package mobile

import (
	"testing"

	_ "github.com/Jaeel/fyne/test"
)

func TestDevice_HideVirtualKeyboard_BeforeRun(t *testing.T) {
	hideVirtualKeyboard() // should not crash!
}
