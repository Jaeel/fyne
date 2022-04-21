// +build !ios,!android

package mobile

import (
	"io"

	"github.com/Jaleel/fyne"
	intRepo "github.com/Jaleel/fyne/internal/repository"
	"github.com/Jaleel/fyne/storage/repository"
)

func existsURI(fyne.URI) (bool, error) {
	// no-op as we use the internal FileRepository
	return false, nil
}

func nativeFileOpen(*fileOpen) (io.ReadCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func nativeFileSave(*fileSave) (io.WriteCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func registerRepository(d *mobileDriver) {
	repository.Register("file", intRepo.NewFileRepository())
}
