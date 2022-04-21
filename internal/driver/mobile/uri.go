// +build !android

package mobile

import (
	"github.com/Jaeel/fyne"
	"github.com/Jaeel/fyne/storage"
)

func nativeURI(uri string) fyne.URI {
	return storage.NewURI(uri)
}
