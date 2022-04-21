package playground

import (
	"github.com/Jaleel/fyne/driver/software"
	"github.com/Jaleel/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
