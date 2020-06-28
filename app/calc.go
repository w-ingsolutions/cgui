package calc

import (
	"gioui.org/layout"
)

func (w *WingCal) GlavniEkran() func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			//layout.Rigid(func(gtx C) D {
			//	return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["DarkGrayI"]).Layout(&gtx, layout.Center, header(w))
			//}),
			layout.Flexed(1, func(gtx C) D {
				//layout.Rigid(func(gtx C) D {
				return w.UI.BezMargine.Layout(gtx, w.strana())
			}))
		//layout.Rigid(func() {
		//	//w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
		//	material.H3("Footer").Layout(w.Context)
		//})
	}
}

func (w *WingCal) strana() (s func(gtx C) D) {
	switch w.Strana {
	case "materijal":
		//return material(w)
	case "kalkulator":
		s = w.kalkulator()
	case "projekat":
		//return projekat(w)
	}
	return
}
