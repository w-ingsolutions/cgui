package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"

	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) PrikazaniElementIzgled() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.Tema.WingUIcontainer(1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Flexed(1, w.PrikazaniElementOpis()),
					layout.Rigid(w.PrikazaniElementSuma()))
			})
		})
	}
}

func (w *WingCal) PrikazaniElementStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = width
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(0.4, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Naziv", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Potrosnja", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Koeficijent", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Merodavna potrosnja", w.Cyr)).Layout(gtx)
			}),
			layout.Flexed(0.15, func(gtx C) D {
				return material.Caption(w.UI.Tema.T, latcyr.C("Cena materijala", w.Cyr)).Layout(gtx)
			}))
	}
}

func (w *WingCal) PrikazaniElementDugmeDodaj(sumaCena float64) func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, dodajDugme, latcyr.C("DODAJ", w.Cyr))
		//btn.FullWidth = true
		//btn.FullHeight = true
		btn.Background = gelook.HexARGB(w.UI.Tema.Colors["Secondary"])
		var varijacijaRada int

		for dodajDugme.Clicked() {
			if kolicina.Value > 0 {
				for _, s := range w.Suma.ElementiPrikaz {
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
			}
			w.NeopodanMaterijal()
			w.SumaRacunica()
			w.Strana = "suma"
		}
		return btn.Layout(gtx)
	}
}
func (w *WingCal) SumaElementiPrikaz() {
	w.Suma.ElementiPrikaz = nil
	for _, e := range w.Suma.Elementi {
		w.Suma.ElementiPrikaz = append(w.Suma.ElementiPrikaz, e)

	}

}
func (w *WingCal) PrikazaniElementOpis() func(gtx C) D {
	return func(gtx C) D {
		//return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["LightGray"]).Layout(gtx, layout.NW, func(gtx C) D {
		//	//sumaCena := 111.33
		//	//width := gtx.Constraints.Max.X
		//	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		//layout.Rigid(func(gtx C) D {
		//	return material.Body1(w.UI.Tema.T, latcyr.C(w.PrikazaniElement.Opis, w.Cyr)).Layout(gtx)
		//}),
		//layout.Rigid(func(gtx C) D { return w.UI.Tema.WingUIline(gtx, 0, 0, 16, w.UI.Tema.Colors["Gray"]) }),
		//layout.Rigid(func(gtx C) D {
		//	return material.Caption(w.UI.Tema.T, latcyr.C(w.PrikazaniElement.Obracun, w.Cyr)).Layout(gtx)
		//}),
		//layout.Rigid(func(gtx C) D { return w.UI.Tema.WingUIline(gtx, 0, 0, 8, w.UI.Tema.Colors["Gray"]) }),
		//layout.Rigid(func(gtx C) D { return w.UI.Tema.WingUIline(gtx, 0, 0, 32, w.UI.Tema.Colors["Gray"]) }),
		//layout.Rigid(func(gtx C) D {
		//	return material.H6(w.UI.Tema.T, latcyr.C("Neophodan materijal za izvrsenje radova", w.Cyr)).Layout(gtx)
		//}),
		//layout.Rigid(w.PrikazaniElementStavkeMaterijala(width)),
		//layout.Rigid(w.RadNeophodanMaterijal(neophodanMaterijalList)),
		//		)
		//	})
		//}
		widgets := []layout.Widget{
			material.H5(w.UI.Tema.T, fmt.Sprint(w.Podvrsta)+"."+fmt.Sprint(w.Roditelj)+"."+fmt.Sprint(w.PrikazaniElement.Id)+" "+latcyr.C(w.PrikazaniElement.Naziv, w.Cyr)).Layout,
			material.Body1(w.UI.Tema.T, latcyr.C(w.PrikazaniElement.Opis, w.Cyr)).Layout,
			material.Caption(w.UI.Tema.T, latcyr.C(w.PrikazaniElement.Obracun, w.Cyr)).Layout,
			material.H6(w.UI.Tema.T, latcyr.C("Neophodan materijal za izvrsenje radova", w.Cyr)).Layout,
			w.PrikazaniElementStavkeMaterijala(),
			w.RadNeophodanMaterijal(neophodanMaterijalList),
		}
		return elementOpis.Layout(gtx, len(widgets), func(gtx C, i int) D {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
		})
	}
}

func (w *WingCal) PrikazaniElementSuma() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.Tema.WingUIcontainer(0, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.NW, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			sumaCena := float64(kolicina.Value) * w.PrikazaniElement.Cena
			return layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.H6(w.UI.Tema.T, latcyr.C("Cena:", w.Cyr)+fmt.Sprint(w.PrikazaniElement.Cena)).Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx C) D { return w.UI.Tema.WingUIline(gtx, 0, 0, 0, w.UI.Tema.Colors["Gray"]) }),
						layout.Rigid(func(gtx C) D {
							return w.UI.Tema.WingUIcontainer(8, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.H6(w.UI.Tema.T, latcyr.C("Suma:", w.Cyr)+fmt.Sprintf("%.2f", w.PrikazaniElement.Cena*float64(kolicina.Value))).Layout(gtx)

							})
						}),
					)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx,
						layout.Rigid(w.UI.Tema.WingUIcounter(kolicina, func() {}).Layout(kolicina, gtx, w.UI.Tema.T, latcyr.C("KOLIČINA", w.Cyr), fmt.Sprint(kolicina.Value))),
						layout.Rigid(w.PrikazaniElementDugmeDodaj(sumaCena)),
					)
				}))
		})
	}
}
