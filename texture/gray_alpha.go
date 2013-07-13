package texture

import (
  "github.com/runningwild/memory"
  "image"
  "image/color"
)

// NOTE: All of this code is basically ripped from the Go source, it's just
// been modified to include an alpha value

type alpha uint8

func (a alpha) RGBA() (uint32, uint32, uint32, uint32) {
  return 0, 0, 0, uint32(a) << 8
}

var GrayAlphaModel alphaModel

type alphaModel struct{}

func (am alphaModel) Convert(c color.Color) color.Color {
  // r, g, b, a := c.RGBA()
  // return alpha{ (r + g + b) / 3, a }
  r, _, _, _ := c.RGBA()
  return am.baseConvert(r)
}
func (am alphaModel) baseConvert(a uint32) color.Color {
  return alpha(a)
}

type GrayAlpha struct {
  // Pix holds the image's pixels, as gray values. The pixel at
  // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
  Pix []uint8
  // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
  Stride int
  // Rect is the image's bounds.
  Rect image.Rectangle
}

func (p *GrayAlpha) ColorModel() color.Model { return GrayAlphaModel }

func (p *GrayAlpha) Bounds() image.Rectangle { return p.Rect }

func (p *GrayAlpha) At(x, y int) color.Color {
  if !(image.Point{x, y}.In(p.Rect)) {
    return alpha(0)
  }
  i := p.PixOffset(x, y)
  return GrayAlphaModel.baseConvert(uint32(p.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayAlpha) PixOffset(x, y int) int {
  return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *GrayAlpha) Set(x, y int, c color.Color) {
  if !(image.Point{x, y}.In(p.Rect)) {
    return
  }
  i := p.PixOffset(x, y)
  r, _, _, _ := c.RGBA()
  p.Pix[i] = byte(r >> 8)
}
func NewGrayAlpha(r image.Rectangle) *GrayAlpha {
  var dx, dy int
  dx = r.Dx()
  // if dx%2 == 1 {
  //   dx++
  // }
  dy = r.Dy()
  return &GrayAlpha{
    Pix:    memory.GetBlock(dx * dy),
    Stride: dx,
    Rect:   r,
  }
}
