// +build android ios mobile

package cache

import "github.com/Jaeel/fyne/internal/driver/mobile/gl"

// TextureType represents an uploaded GL texture
type TextureType = gl.Texture

var noTexture = gl.Texture{0}

type textureInfo struct {
	textureCacheBase
	texture TextureType
}
