package software

import (
	"github.com/Jaeel/fyne/internal/painter/software"
	"github.com/Jaeel/fyne/test"
)

// NewCanvas creates a new canvas in memory that can render without hardware support
func NewCanvas() test.WindowlessCanvas {
	return test.NewCanvasWithPainter(software.NewPainter())
}
