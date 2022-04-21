package app_test

import (
	"testing"

	"github.com/Jaleel/fyne/internal/app"
	"github.com/Jaleel/fyne/test"
)

func TestApplySettings_BeforeContentSet(t *testing.T) {
	a := test.NewApp()
	_ = a.NewWindow("NoContent")

	app.ApplySettings(a.Settings(), a)
}
