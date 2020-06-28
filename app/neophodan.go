package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) RadNeophodanMaterijal(l *layout.List) func(gtx C) D {
	return func(gtx C) D {
		//var materijal model.WingNeophodanMaterijal
		nm := w.PrikazaniElement.NeophodanMaterijal
		width := gtx.Constraints.Max.X
		return l.Layout(gtx, len(nm), func(gtx C, i int) D {
			materijal := nm[i]
			id := materijal.Id - 1
			materijal.Koeficijent = materijal.Koeficijent
			materijal.Materijal = *w.Materijal[id]
			if materijal.Koeficijent > 0 {
				materijal.Kolicina = materijal.Materijal.Potrosnja * float64(kolicina.Value) * materijal.Koeficijent
			}
			materijal.UkupnaCena = materijal.Materijal.Cena * float64(materijal.Kolicina)
			materijal.UkupnoPakovanja = int(materijal.Kolicina / float64(materijal.Materijal.Pakovanje))

			gtx.Constraints.Min.X = width

			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(gtx,
						layout.Flexed(0.4, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, latcyr.C(materijal.Materijal.Naziv, w.Cyr)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprintf("%.2f", materijal.Materijal.Potrosnja)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprint(materijal.Koeficijent)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprintf("%.2f", materijal.Kolicina)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprintf("%.2f", materijal.UkupnaCena)).Layout(gtx)
						}),
					)
				}),
				layout.Rigid(w.UI.Tema.WingUIline(&gtx, 0, 0, 1, w.UI.Tema.Colors["Gray"])),
			)
		})
	}
}

func (w *WingCal) UkupanNeophodanMaterijal(l *layout.List) func(gtx C) D {
	return func(gtx C) D {

		//fmt.Println(":::::UkupanNeophodanMaterijal", w.Suma.UkupanNeophodanMaterijal)
		//var materijal model.WingNeophodanMaterijal
		width := gtx.Constraints.Max.X
		return l.Layout(gtx, len(w.Suma.UkupanNeophodanMaterijal), func(gtx C, i int) D {

			materijal := w.Suma.UkupanNeophodanMaterijalPrikaz[i]
			//materijal.Materijal = *w.Materijal[materijal.Id-1]
			//fmt.Println(":::::NazivNazivNaziv", materijal.Materijal.Naziv)
			//fmt.Println(":::::IDidididi111", materijal.Id)
			//fmt.Println(":::::IDidididi", materijal.Materijal.Id)
			gtx.Constraints.Min.X = width

			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(gtx,
						layout.Flexed(0.4, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, latcyr.C(materijal.Materijal.Naziv, w.Cyr)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprint(materijal.Materijal.Cena)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprintf("%.2f", materijal.Kolicina)).Layout(gtx)
						}),
						layout.Flexed(0.15, func(gtx C) D {
							return material.Body1(w.UI.Tema.T, fmt.Sprintf("%.2f", materijal.UkupnaCena)).Layout(gtx)
						}),
					)
				}),
				layout.Rigid(w.UI.Tema.WingUIline(&gtx, 0, 0, 1, w.UI.Tema.Colors["Gray"])),
			)
		})
	}
}

func (w *WingCal) SumaStavkeMaterijala(width int) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = width
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(0.4, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Naziv", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Pojedinacna cena", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Kolicina", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Ukupna cena", w.Cyr)).Layout(gtx)
			}))
	}
}
