package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

var (
	tabelaSuma = map[int]int{}
)

func (w *WingCal) SumaIzgled() func(gtx C) D {
	return func(gtx C) D {
		//w.Tema.DuoUIcontainer(0, w.Tema.Colors["LightGrayI"]).Layout(w.Context, layout.NW, func() {
		width := gtx.Constraints.Max.X
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.5, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						ukupan := w.UI.Tema.WingUIcontainer(16, w.UI.Tema.Colors["Primary"])
						ukupan.FullWidth = true
						return ukupan.Layout(&gtx, layout.W, func(gtx C) D {
							suma := material.H6(w.UI.Tema.T, latcyr.C("Ukupna cena radova", w.Cyr))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})

					}),
					layout.Flexed(1, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
									return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
										layout.Flexed(1, func(gtx C) D {
											return material.Caption(w.UI.Tema.T, "Naziv").Layout(gtx)
										}),
										layout.Rigid(w.cell(gtx, 0, "Pojedinačna cena")),
										layout.Rigid(w.cell(gtx, 1, "Količina")),
										layout.Rigid(w.cell(gtx, 2, "Cena")),
										layout.Rigid(w.cell(gtx, 3, "briši")))
								})
							}),
							layout.Rigid(func(gtx C) D {
								return sumList.Layout(gtx, len(w.Suma.Elementi), func(gtx C, i int) D {
									element := w.Suma.Elementi[i]
									return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
										return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
											layout.Flexed(1, func(gtx C) D {
												return material.Body1(w.UI.Tema.T, element.Element.Naziv).Layout(gtx)
											}),
											layout.Rigid(w.cell(gtx, 0, fmt.Sprint(element.Element.Cena))),
											layout.Rigid(w.cell(gtx, 1, fmt.Sprint(element.Kolicina))),
											layout.Rigid(w.cell(gtx, 2, fmt.Sprintf("%.2f", element.SumaCena))),
											layout.Rigid(func(gtx C) D {
												btn := material.Button(w.UI.Tema.T, element.DugmeBrisanje, latcyr.C("BRIŠI", w.Cyr)+fmt.Sprint(i))
												btn.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
												btn.TextSize = unit.Dp(12)
												btn.Color = gelook.HexARGB(w.UI.Tema.Colors["Gray"])
												btn.Background = gelook.HexARGB(w.UI.Tema.Colors["yellow"])
												for element.DugmeBrisanje.Clicked() {
													fmt.Println("iii", i)

													//fmt.Println("w.Suma.ElementiPREEEE", w.Suma.Elementi)
													w.Suma.Elementi = append(w.Suma.Elementi[:i], w.Suma.Elementi[i+1:]...)
													//w.Suma.Elementi[fmt.Sprint(i)] =  model.WingIzabraniElement{}
													tabelaSuma = map[int]int{}
													w.NeopodanMaterijal()
													w.SumaRacunica()
												}
												return btn.Layout(gtx)
												//w.tabela(3, w.Context.Dimensions.Size.X)
											}))
									})
								})
							}))
					}),
					layout.Rigid(func(gtx C) D {
						gtx.Constraints.Min.X = width
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["LightGrayII"]).Layout(&gtx, layout.SE, func(gtx C) D {
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
						ukupan := w.UI.Tema.WingUIcontainer(16, w.UI.Tema.Colors["Primary"])
						ukupan.FullWidth = true
						return ukupan.Layout(&gtx, layout.W, func(gtx C) D {
							suma := material.H6(w.UI.Tema.T, latcyr.C("Ukupan neophodni materijal", w.Cyr))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})
					}),
					layout.Rigid(w.SumaStavkeMaterijala(width)),
					layout.Rigid(w.UI.Tema.WingUIline(&gtx, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(1, w.UkupanNeophodanMaterijal(ukupanNeophodanMaterijalList)),

					layout.Rigid(func(gtx C) D {
						gtx.Constraints.Min.X = width
						return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["LightGrayII"]).Layout(&gtx, layout.SE, func(gtx C) D {
							suma := material.Body2(w.UI.Tema.T, latcyr.C("Suma materijal: ", w.Cyr)+fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})
					}),

					layout.Rigid(w.Stampa()))
			}))
		//)
	}
	//return  w.UI.D
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

func (w *WingCal) tabela(d D, colona, width int) {
	if width > tabelaSuma[colona] {
		tabelaSuma[colona] = width
	}
	d.Size.X = tabelaSuma[colona]
}

func (w *WingCal) cell(gtx layout.Context, colona int, tekst string) func(gtx C) D {
	return func(gtx C) D {
		var d D
		w.tabela(d, colona, d.Size.X)
		return layout.Inset{
			Top:    unit.Dp(0),
			Right:  unit.Dp(4),
			Bottom: unit.Dp(0),
			Left:   unit.Dp(4),
		}.Layout(gtx, func(gtx C) D {
			return material.Caption(w.UI.Tema.T, tekst).Layout(gtx)
		})
	}
}

func (w *WingCal) SumaRacunica() {
	s := 0.0
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena
	}
	w.Suma.SumaCena = s
}
