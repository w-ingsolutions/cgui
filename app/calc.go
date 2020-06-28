package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func (w *WingCal) GlavniEkran() layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(w.UI.Context,
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
		case d > 1000:
			w.UI.Device = "h"
		case d < 1000:
			w.UI.Device = "p"
		}
		var s D

		switch w.UI.Device {
		case "p":
			fmt.Println("pppp")

			switch w.Strana {
			case "materijal":
				s = w.UI.BezMargine.Layout(gtx, w.MaterijalStrana())
			case "izbornik":
				s = w.UI.BezMargine.Layout(gtx, w.IzbornikRadovaStrana())
			case "suma":
				s = w.UI.BezMargine.Layout(gtx, w.SumaStrana())
			}
		case "h":
			fmt.Println("hhhh")

			s = layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.UI.Context,
				layout.Flexed(0.2, func(gtx C) D {
					return w.UI.BezMargine.Layout(gtx, w.IzbornikRadovaStrana())
				}),
				layout.Flexed(0.2, func(gtx C) D {
					return w.UI.BezMargine.Layout(gtx, w.MaterijalStrana())
				}),
				//layout.Flexed(0.2, func(gtx C) D {
				//	//return w.UI.BezMargine.Layout(gtx, w.kalkulator())
				//}),
				layout.Flexed(0.2, func(gtx C) D {
					return w.UI.BezMargine.Layout(gtx, w.SumaStrana())
					//}),
					//layout.Flexed(0.2, func(gtx C) D {
					//return w.UI.BezMargine.Layout(gtx, w.kalkulator())
				}))
		}
		return s

		//return layout.Flex{
		//	Axis: layout.Horizontal,
		//}.Layout(w.UI.Context,
		//	layout.Flexed(1, func(gtx C) D {
		//		return w.UI.BezMargine.Layout(gtx, s)
		//	}))

	}
}
