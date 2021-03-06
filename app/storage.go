package app

import (
	"os"

	"github.com/Jaeel/fyne"
	"github.com/Jaeel/fyne/internal"
	"github.com/Jaeel/fyne/storage"
)

type store struct {
	*internal.Docs
	a *fyneApp
}

func (s *store) RootURI() fyne.URI {
	if s.a.uniqueID == "" {
		fyne.LogError("Storage API requires a unique ID, use app.NewWithID()", nil)
		return storage.NewURI("file://" + os.TempDir())
	}

	return storage.NewURI("file://" + s.a.storageRoot())
}

func (s *store) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
