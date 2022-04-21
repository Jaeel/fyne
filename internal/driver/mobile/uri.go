// +build !android

package mobile

import (
	"github.com/Jaleel/fyne"
	"github.com/Jaleel/fyne/storage"
)

func nativeURI(uri string) fyne.URI {
	return storage.NewURI(uri)
}
