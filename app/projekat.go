package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
)

func (w *WingCal) ProjekatStrana() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						return projekatList.Layout(gtx, len(w.projekat()), func(gtx C, i int) D {
							return layout.Flex{}.Layout(gtx,
								layout.Flexed(0.5, w.projekat()[i]))
						})
					}),
					layout.Rigid(w.Stampa()))
			})
		})
	}
}

func (w *WingCal) projekat() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.4, func(gtx C) D {
						return projektantiList.Layout(gtx, len(w.Lica.Projektanti), func(gtx C, i int) D {
							return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
								return material.RadioButton(w.UI.Tema.T, projektantIzbor, fmt.Sprint(w.Lica.Projektanti[i].Id), w.Lica.Projektanti[i].Ime+" "+w.Lica.Projektanti[i].Prezime).Layout(gtx)
							})
						})
					}),
					layout.Flexed(0.6, func(gtx C) D {
						return layout.Flex{}.Layout(gtx,
							layout.Flexed(0.4, func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Ime")+":"+w.text(projektantIzbor.Value)).Layout(gtx)
							}),
							layout.Flexed(0.4, func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Ime")+":"+w.text(projektantIzbor.Value)).Layout(gtx)
							}))
					}))
			})
		},
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				return klijentiList.Layout(gtx, len(w.Lica.Klijenti), func(gtx C, i int) D {
					return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
						return material.RadioButton(w.UI.Tema.T, klijentiIzbor, fmt.Sprint(w.Lica.Klijenti[i].Id), w.Lica.Klijenti[i].Naziv).Layout(gtx)
					})
				})
			})
		},
	}
}
