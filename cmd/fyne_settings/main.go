package main

import (
	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/app"
	"github.com/Jaleel/fyne/cmd/fyne_settings/settings"
	"github.com/Jaleel/fyne/container"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := container.NewAppTabs(
		&container.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(480, 480))
	w.ShowAndRun()
}
