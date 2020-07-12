package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
)

var (
	tabelaSuma = map[int]int{}
)

func (w *WingCal) SumaStranaPrazno() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.W, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			prazno := material.H3(w.UI.Tema.T, w.text("Izaberite radove"))
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
						return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(w.UI.Tema.T, w.text("Ukupna cena radova"))
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})

					}),
					layout.Rigid(func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx, w.SumaRadoviStavke())
							}))
					}),
					layout.Rigid(helper.DuoUIline(false, 4, 4, 2, w.UI.Tema.Colors["Primary"])),
					layout.Flexed(1, w.SumaElementi(w.Suma.Elementi)))
			}),
			layout.Rigid(func(gtx C) D {
				return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, func(gtx C) D {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					suma := material.H5(w.UI.Tema.T, w.text("Suma: ")+fmt.Sprintf("%.2f", w.Suma.SumaCena))
					suma.Alignment = text.End
					return suma.Layout(gtx)
				})
			}),
			layout.Flexed(0.5, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(w.UI.Tema.T, w.text("Ukupan neophoddddni materijal"))
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})
					}),
					layout.Rigid(w.SumaStavkeMaterijala()),
					layout.Rigid(helper.DuoUIline(false, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(1, w.UkupanNeophodanMaterijal(w.Suma.NeophodanMaterijal)),

					layout.Rigid(func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.Body2(w.UI.Tema.T, w.text("Sumaaaaa materijal: ")+fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal))
							suma.Alignment = text.End
							return suma.Layout(gtx)
						})
					}))

				//layout.Rigid(w.Stampa()))
			}))
		//)
		//return  w.UI.D
	}
}

func (w *WingCal) SumaElementi(el []*model.WingIzabraniElement) func(gtx C) D {
	return func(gtx C) D {
		return sumList.Layout(gtx, len(el), func(gtx C, i int) D {
			element := el[i]
			return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
					layout.Flexed(0.6, w.cell(text.Start, w.text(element.Element.Naziv))),
					layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(0.1, w.cell(text.Middle, fmt.Sprint(element.Element.Cena))),
					layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(0.1, w.cell(text.Middle, fmt.Sprint(element.Kolicina))),
					layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
					layout.Flexed(0.1, w.cell(text.Middle, fmt.Sprintf("%.2f", element.SumaCena))),
					layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
					layout.Rigid(w.brisi(element, i)))
			})
		})
	}
}
