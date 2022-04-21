package software

import (
	"github.com/Jaleel/fyne/internal/painter/software"
	"github.com/Jaleel/fyne/test"
)

// NewCanvas creates a new canvas in memory that can render without hardware support
func NewCanvas() test.WindowlessCanvas {
	return test.NewCanvasWithPainter(software.NewPainter())
}
