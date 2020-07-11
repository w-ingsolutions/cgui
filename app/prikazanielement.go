package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) PrikazaniElementDugmeDodaj(sumaCena float64) func(gtx C) D {
	return func(gtx C) D {
		return w.UI.BezMargine.Layout(gtx, func(gtx C) D {
			//fmt.Println("maxOPED:", gtx.Constraints.Max.X)
			//fmt.Println("minOPED:", gtx.Constraints.Min.X)
			btn := material.Button(w.UI.Tema.T, dodajDugme, latcyr.C("DODAJ", w.Podesavanja.Cyr))
			//btn.FullWidth = true
			//btn.FullHeight = true
			btn.CornerRadius = unit.Dp(0)
			btn.TextSize = unit.Dp(14)
			btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
			btn.Background = helper.HexARGB("ffb8df42")
			btn.Color = helper.HexARGB(w.UI.Tema.Colors["Dark"])
			var varijacijaRada int

			for dodajDugme.Clicked() {
				fmt.Println("kolicina.Value", kolicina.Value)
				if kolicina.Value > 0 {
					for _, s := range w.Suma.Elementi {
						if s.Element.Id == w.PrikazaniElement.Id {
							varijacijaRada = varijacijaRada + 1
							fmt.Println("varijacijaRada:", varijacijaRada)
						}
						fmt.Println("elem:", s.Element.Id)
					}
					fmt.Println("varijacijaRadaIIIIIIIIIIIII:", varijacijaRada)
					suma := model.WingIzabraniElement{
						Sifra:         fmt.Sprint(w.Podvrsta) + "." + fmt.Sprint(w.Roditelj) + "." + fmt.Sprint(w.PrikazaniElement.Id) + "." + fmt.Sprint(varijacijaRada+1),
						Kolicina:      kolicina.Value,
						SumaCena:      sumaCena,
						Element:       *w.PrikazaniElement,
						DugmeBrisanje: new(widget.Clickable),
					}
					w.Suma.Elementi = append(w.Suma.Elementi, &suma)
					w.NeopodanMaterijal()
					w.SumaRacunica()
					w.Strana = "sumaRadovi"
				}
			}
			return btn.Layout(gtx)
		})
	}
}
func (w *WingCal) SumaElementiPrikaz() {
	//w.Suma.ElementiPrikaz = nil
	//for _, e := range w.Suma.Elementi {
	//	w.Suma.ElementiPrikaz = append(w.Suma.Elementi, e)
	//
	//}

}
func (w *WingCal) PrikazaniElementIzgled() func(gtx C) D {
	return func(gtx C) D {
		neophodanNaslov := material.H6(w.UI.Tema.T, w.text("Neophodan materijal za izvrsenje radova"))
		neophodanNaslov.Color = helper.HexARGB(w.UI.Tema.Colors["Primary"])
		widgets := []layout.Widget{
			material.H5(w.UI.Tema.T, fmt.Sprint(w.Podvrsta)+"."+fmt.Sprint(w.Roditelj)+"."+fmt.Sprint(w.PrikazaniElement.Id)+" "+w.text(w.PrikazaniElement.Naziv)).Layout,
			material.Body1(w.UI.Tema.T, w.text(w.PrikazaniElement.Opis)).Layout,
			material.Caption(w.UI.Tema.T, w.text(w.PrikazaniElement.Obracun)).Layout,
			neophodanNaslov.Layout,
			helper.DuoUIline(false, 4, 0, 4, w.UI.Tema.Colors["Secondary"]),
			w.PrikazaniElementStavkeMaterijala(),
			helper.DuoUIline(false, 4, 0, 2, w.UI.Tema.Colors["Primary"]),
			w.RadNeophodanMaterijal(neophodanMaterijalList),
		}
		return elementOpis.Layout(gtx, len(widgets), func(gtx C, i int) D {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
		})
	}
}

func (w *WingCal) PrikazaniElementSuma() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.NW, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			sumaCena := float64(kolicina.Value) * w.PrikazaniElement.Cena
			return layout.Flex{
				Axis:      layout.Horizontal,
				Spacing:   layout.SpaceBetween,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Flexed(0.5, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.Body2(w.UI.Tema.T, w.text("Cena:")+fmt.Sprint(w.PrikazaniElement.Cena)).Layout(gtx)
							})
						}),
						layout.Rigid(helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Dark"])),
						layout.Rigid(func(gtx C) D {
							return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.Body2(w.UI.Tema.T, latcyr.C("Suma:", w.Podesavanja.Cyr)+fmt.Sprintf("%.2f", prikazaniElementSumaCena)).Layout(gtx)

							})
						}),
					)
				}),
				layout.Flexed(0.5, func(gtx C) D {
					return w.UI.BezMargine.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Spacing:   layout.SpaceBetween,
							Alignment: layout.Middle,
						}.Layout(gtx,
							layout.Rigid(counter.DuoUIcounterSt(w.UI.Tema, kolicina, w.PrikazaniElementSumaRacunica()).Layout(gtx, w.UI.Tema.T, latcyr.C("KOLIČINA", w.Podesavanja.Cyr), fmt.Sprint(kolicina.Value))),
							layout.Rigid(w.PrikazaniElementDugmeDodaj(sumaCena)),
						)
					})
				}))
		})
	}
}
