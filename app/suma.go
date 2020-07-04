package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

var (
	tabelaSuma = map[int]int{}
)

func (w *WingCal) SumaStrana() func(gtx C) D {
	return func(gtx C) D {
		s := w.SumaStranaPrazno()
		if w.Suma.ElementiPrikaz != nil {
			s = w.SumaStranaUnutra()
		}
		return w.UI.Tema.WingUIcontainer(1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, s)
		})
	}
}

func (w *WingCal) SumaStranaPrazno() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.Tema.WingUIcontainer(4, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			prazno := material.H3(w.UI.Tema.T, latcyr.C("Izaberite radove", w.Cyr))
			prazno.Alignment = text.Middle
			return prazno.Layout(gtx)
		})
	}
}

func (w *WingCal) SumaStranaUnutra() func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.5, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return w.UI.Tema.WingUIcontainer(4, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(w.UI.Tema.T, latcyr.C("Ukupna cena radova", w.Cyr))
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})

					}),
					layout.Flexed(1, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
									return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
										layout.Flexed(0.6, w.cell(latcyr.C("Naziv", w.Cyr))),
										layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
										layout.Flexed(0.15, w.cell(latcyr.C("Pojedinačna cena", w.Cyr))),
										layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
										layout.Flexed(0.1, w.cell(latcyr.C("Količina", w.Cyr))),
										layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
										layout.Flexed(0.1, w.cell(latcyr.C("Cena", w.Cyr))),
										layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
										layout.Flexed(0.05, w.cell(latcyr.C("briši", w.Cyr))))
								})
							}),
							layout.Rigid(func(gtx C) D {
								return sumList.Layout(gtx, len(w.Suma.ElementiPrikaz), func(gtx C, i int) D {
									element := w.Suma.Elementi[i]
									return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
										return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
											layout.Flexed(0.6, w.cell(latcyr.C(element.Element.Naziv, w.Cyr))),
											layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
											layout.Flexed(0.1, w.cell(fmt.Sprint(element.Element.Cena))),
											layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
											layout.Flexed(0.1, w.cell(fmt.Sprint(element.Kolicina))),
											layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
											layout.Flexed(0.1, w.cell(fmt.Sprintf("%.2f", element.SumaCena))),
											layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
											layout.Rigid(w.brisi(element, i)))
									})
								})
							}))
					}),
					layout.Rigid(func(gtx C) D {
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H5(w.UI.Tema.T, latcyr.C("Suma: ", w.Cyr)+fmt.Sprintf("%.2f", w.Suma.SumaCena))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})
					}),
				)
			}),

			layout.Flexed(0.5, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return w.UI.Tema.WingUIcontainer(4, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(w.UI.Tema.T, latcyr.C("Ukupan neophodni materijal", w.Cyr))
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})
					}),
					layout.Rigid(w.SumaStavkeMaterijala()),
					layout.Rigid(w.UI.Tema.WingUIline(false, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(1, w.UkupanNeophodanMaterijal(ukupanNeophodanMaterijalList)),

					layout.Rigid(func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.Body2(w.UI.Tema.T, latcyr.C("Suma materijal: ", w.Cyr)+fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})
					}),

					layout.Rigid(w.Stampa()))
			}))
		//)
		//return  w.UI.D
	}
}

func (w *WingCal) NeopodanMaterijal() {
	u := make(map[int]model.WingNeophodanMaterijal)
	unm := make(map[int]model.WingNeophodanMaterijal)
	sumaCena := 0.0
	for _, e := range w.Suma.Elementi {
		for _, n := range e.Element.NeophodanMaterijal {
			uu := model.WingNeophodanMaterijal{
				Id:        n.Id,
				Materijal: *w.Materijal[n.Id-1],
			}
			k := uu.Materijal.Potrosnja * float64(e.Kolicina) * n.Koeficijent
			uu.Kolicina = u[n.Id].Kolicina + k
			ukupnaCena := uu.Kolicina * uu.Materijal.Cena
			uu.UkupnaCena = ukupnaCena
			uu.UkupnoPakovanja = int(k / float64(uu.Materijal.Pakovanje))
			u[n.Id] = uu
			sumaCena = sumaCena + ukupnaCena
		}
	}
	w.Suma.UkupanNeophodanMaterijal = u
	w.Suma.SumaCenaMaterijal = sumaCena
	i := 0
	for _, uuu := range u {
		unm[i] = uuu
		i++
	}
	w.Suma.UkupanNeophodanMaterijalPrikaz = unm
}

func (w *WingCal) SumaRacunica() {
	w.SumaElementiPrikaz()
	s := 0.0
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena
	}
	w.Suma.SumaCena = s
}
