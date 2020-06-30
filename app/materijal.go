package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func (w *WingCal) MaterijalStrana() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.SaMarginom.Layout(gtx, func(gtx C) D {
			return materijalList.Layout(gtx, len(w.Materijal), func(gtx C, i int) D {
				m := w.Materijal[i]
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Flexed(0.1, func(gtx C) D {
								return material.H6(w.UI.Tema.T, fmt.Sprint(m.Id)).Layout(gtx)
							}),
							layout.Flexed(0.4, func(gtx C) D {
								return material.H6(w.UI.Tema.T, m.Naziv).Layout(gtx)
							}),
							layout.Flexed(0.2, func(gtx C) D {
								return material.Caption(w.UI.Tema.T, m.Obracun).Layout(gtx)
							}),
							layout.Flexed(0.1, func(gtx C) D {
								return material.Body2(w.UI.Tema.T, fmt.Sprint(m.Pakovanje)).Layout(gtx)
							}),
							layout.Flexed(0.1, func(gtx C) D {
								return material.Body1(w.UI.Tema.T, m.Jedinica).Layout(gtx)
							}),
							layout.Flexed(0.1, func(gtx C) D {
								return material.H6(w.UI.Tema.T, fmt.Sprint(m.Cena)).Layout(gtx)
							}),
						)
					}),

					layout.Rigid(func(gtx C) D {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx)//layout.Flexed(0.4, func(gtx C) D {
						//	return material.Body1(w.UI.Tema.T, m.NacinRada).Layout(gtx)
						//}),

						//layout.Flexed(0.2, func(gtx C) D {
						//	return material.Caption(w.UI.Tema.T, m.OsobineNamena).Layout(gtx)
						//}),

						//layout.Flexed(0.1, func(gtx C) D {
						//	return material.H6(w.UI.Tema.T, fmt.Sprint(m.Potrosnja)).Layout(gtx)
						//}),
						//layout.Flexed(0.1, func(gtx C) D {
						//	return material.H6(w.UI.Tema.T, m.JedinicaPotrosnje).Layout(gtx)
						//}),
						//layout.Flexed(0.2, func(gtx C) D {
						//	return material.H6(w.UI.Tema.T, m.RokUpotrebe).Layout(gtx)
						//}),

					}),
					layout.Rigid(func(gtx C) D { return w.UI.Tema.WingUIline(gtx, 1, 1, 1, w.UI.Tema.Colors["Dark"]) }),
				)
			})
		})
	}
}
