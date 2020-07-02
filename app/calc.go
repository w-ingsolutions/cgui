package calc

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func (w *WingCal) GlavniEkran() layout.Dimensions {
	var top layout.Dimensions
	if w.UI.Device == "p" {
		top = w.UI.Tema.WingUIline(w.UI.Context, 32, 0, 0, w.UI.Tema.Colors["Light"])
	}
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(w.UI.Context,

		layout.Rigid(func(gtx C) D {
			return top
		}),

		layout.Rigid(func(gtx C) D {
			return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, header(w))
		}),
		layout.Flexed(1, func(gtx C) D {
			return w.UI.BezMargine.Layout(gtx, w.strana())
		}),
		layout.Rigid(func(gtx C) D {
			//w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
			return material.H3(w.UI.Tema.T, "Footer").Layout(gtx)
		}))
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
			case "izbornik":
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
