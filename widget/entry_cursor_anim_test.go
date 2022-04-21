package widget

import (
	"image/color"
	"runtime"
	"testing"
	"time"

	"github.com/Jaeel/fyne/canvas"
	col "github.com/Jaeel/fyne/internal/color"
	_ "github.com/Jaeel/fyne/test"
	"github.com/Jaeel/fyne/theme"
	"github.com/stretchr/testify/assert"
)

func TestEntryCursorAnim(t *testing.T) {
	cursorOpaque := theme.PrimaryColor()
	r, g, b, _ := col.ToNRGBA(theme.PrimaryColor())
	cursorDim := color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0x16}

	alphaEquals := func(color1, color2 color.Color) bool {
		_, _, _, a1 := col.ToNRGBA(color1)
		_, _, _, a2 := col.ToNRGBA(color2)
		return a1 == a2
	}

	cursor := canvas.NewRectangle(color.Black)
	a := newEntryCursorAnimation(cursor)

	a.start()
	a.anim.Tick(0.0)
	assert.True(t, alphaEquals(cursorDim, a.cursor.FillColor))
	a.anim.Tick(1.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))

	a.interrupt()
	a.anim.Tick(0.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))
	a.anim.Tick(0.5)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))
	a.anim.Tick(1.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))

	a.timeNow = func() time.Time {
		return time.Now().Add(cursorInterruptTime)
	}
	// animation should be restarted inverting the colors
	a.anim.Tick(0.0)
	runtime.Gosched()
	time.Sleep(10 * time.Millisecond) // ensure go routine for restart animation is executed
	a.anim.Tick(0.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))
	a.anim.Tick(1.0)
	assert.True(t, alphaEquals(cursorDim, a.cursor.FillColor))

	a.timeNow = time.Now
	a.interrupt()
	a.anim.Tick(0.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))

	a.timeNow = func() time.Time {
		return time.Now().Add(cursorInterruptTime)
	}
	a.anim.Tick(0.0)
	runtime.Gosched()
	time.Sleep(10 * time.Millisecond) // ensure go routine for restart animation is executed
	a.anim.Tick(0.0)
	assert.True(t, alphaEquals(cursorOpaque, a.cursor.FillColor))
	a.anim.Tick(1.0)
	assert.True(t, alphaEquals(cursorDim, a.cursor.FillColor))

	a.stop()
	assert.Nil(t, a.anim)
}
