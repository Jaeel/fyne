package playground

import (
	"github.com/Jaeel/fyne/driver/software"
	"github.com/Jaeel/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
