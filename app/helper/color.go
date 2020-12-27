package helper

import (
	"fmt"
	"image"
	"image/color"
	
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func HexARGB(s string) (c color.NRGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}

func Rgb(c uint32) color.RGBA {
	return Argb(0xff000000 | c)
}

func Argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

// func Fill(gtx layout.Context, col color.RGBA) layout.Dimensions {
// 	cs := gtx.Constraints
// 	d := cs.Min
// 	dr := f32.Rectangle{
// 		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
// 	}
// 	paint.ColorOp{Color: col}.Add(gtx.Ops)
// 	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
// 	return layout.Dimensions{Size: d}
// }

func Fill(gtx layout.Context, col color.NRGBA, bounds image.Point) {
	clip.UniformRRect(f32.Rectangle{
		Max: f32.Pt(float32(bounds.X), float32(bounds.Y)),
	}, 0).Add(gtx.Ops)
	paint.Fill(gtx.Ops, col)
}
