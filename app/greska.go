package calc

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func (w *WingCal) GreskaEkran() func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(material.H3(w.UI.Tema.T, "Greška u povezivanju sa bazom ").Layout))
	}
}
