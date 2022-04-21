package test

import (
	"os"

	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/internal"
	"github.com/Jaleel/fyne/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewURI("file://" + os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
