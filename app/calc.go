package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func (w *WingCal) GlavniEkran(gtx layout.Context) {
	if w.UI.Device == "p" {
		w.UI.TopSpace = 64
		w.UI.BottomSpace = 64
	} else {
		w.UI.TopSpace = 0
		w.UI.BottomSpace = 0
	}
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(w.UI.Tema.WingUIline(false, 0, 0, w.UI.TopSpace, w.UI.Tema.Colors["Dark"])),
		layout.Flexed(1, func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, header(w))
				}),
				layout.Flexed(1, w.strana()),
				layout.Rigid(func(gtx C) D {
					return material.H6(w.UI.Tema.T, "Footer").Layout(gtx)
				}))

		}),
		layout.Rigid(w.UI.Tema.WingUIline(false, 0, 0, w.UI.BottomSpace, w.UI.Tema.Colors["Dark"])))
}
func (w *WingCal) strana() func(gtx C) D {
	return func(gtx C) D {
		switch d := gtx.Constraints.Max.X; {
		case d > 1100:
			w.UI.Device = "h"
		case d < 1100:
			w.UI.Device = "p"
		}
		var s D

		switch w.UI.Device {
		case "p":
			switch w.Strana {
			case "materijal":
				s = w.UI.SaMarginom.Layout(gtx, w.MaterijalStrana())
			case "radovi":
				s = w.UI.SaMarginom.Layout(gtx, w.IzbornikRadovaStrana())
			case "suma":
				s = w.UI.SaMarginom.Layout(gtx, w.SumaStrana())
			}
		case "h":
			//s = w.UI.SaMarginom.Layout(gtx, w.MaterijalStrana())
			s = w.UI.SaMarginom.Layout(gtx, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Flexed(0.3, func(gtx C) D {
						return w.UI.SaMarginom.Layout(gtx, w.IzbornikRadovaStrana())
						//return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["Light"]).Layout(gtx, layout.N, w.IzbornikRadovaStrana())
					}),
					layout.Flexed(0.4, func(gtx C) D {
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["Light"]).Layout(gtx, layout.N, w.SumaStrana())

					}),
					layout.Flexed(0.3, func(gtx C) D {
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["Light"]).Layout(gtx, layout.N, w.MaterijalStrana())

					}),
				)
			})
		}
		return s
	}

	//return layout.Flex{
	//	Axis: layout.Horizontal,
	//}.Layout(w.UI.Context,
	//	layout.Flexed(1, func(gtx C) D {
	//		return w.UI.BezMargine.Layout(gtx, s)
	//	}))
}

func (w *WingCal) cell(align text.Alignment, tekst string) func(gtx C) D {
	return func(gtx C) D {
		return layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx C) D {
			return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["LightGray"]).Layout(gtx, layout.N, func(gtx C) D {
				return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					//gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
					cell := material.Caption(w.UI.Tema.T, tekst)
					cell.TextSize = unit.Dp(12)
					cell.Alignment = align
					return cell.Layout(gtx)
				})
			})
		})
	}
}
