package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) MaterijalStrana() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.Tema.WingUIcontainer(1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {

				return materijalList.Layout(gtx, len(w.Materijal), func(gtx C, i int) D {
					m := w.Materijal[i]
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
								return layout.Flex{
									Axis:      layout.Horizontal,
									Alignment: layout.Middle,
								}.Layout(gtx,
									layout.Flexed(0.1, w.cell(latcyr.C(fmt.Sprint(m.Id), w.Cyr))),
									layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
									layout.Flexed(0.3, w.cell(latcyr.C(m.Naziv, w.Cyr))),
									layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
									layout.Flexed(0.1, w.cell(latcyr.C(fmt.Sprint(m.Pakovanje), w.Cyr))),
									layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
									layout.Flexed(0.1, w.cell(latcyr.C(m.Jedinica, w.Cyr))),
									layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
									layout.Flexed(0.2, w.cell(latcyr.C(fmt.Sprint(m.Cena), w.Cyr))),
								)
							})
						}),
						//layout.Rigid(func(gtx C) D {
						//	return layout.Flex{
						//		Axis: layout.Horizontal,
						//	}.Layout(gtx) //layout.Flexed(0.4, func(gtx C) D {
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
						//}),
						layout.Rigid(w.UI.Tema.WingUIline(false, 0, 0, 1, w.UI.Tema.Colors["Dark"])),
					)

				})
			})
		})
	}
}
