package test

import (
	"os"

	"github.com/Jaeel/fyne"
	"github.com/Jaeel/fyne/internal"
	"github.com/Jaeel/fyne/storage"
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
