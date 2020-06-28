package calc

import (
	"gioui.org/layout"
)

func (w *WingCal) kalkulator() func(gtx C) D {
	return func(gtx C) D {
		levo := w.IzborVrsteRadova()
		if w.Element {
			levo = w.PrikazaniElementIzgled()
		}
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Flexed(0.5, levo),
			layout.Flexed(0.5, w.SumaIzgled()),
		)
	}
}
