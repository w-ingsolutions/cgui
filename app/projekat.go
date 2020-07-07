package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"strconv"
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
				id, _ := strconv.Atoi(projektantIzbor.Value)
				projektant := w.Lica.Projektanti[id]
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.4, func(gtx C) D {
						return projektantiList.Layout(gtx, len(w.Lica.Projektanti), func(gtx C, i int) D {
							p := w.Lica.Projektanti[i]
							return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
								return material.RadioButton(w.UI.Tema.T, projektantIzbor, fmt.Sprint(p.Id), p.Ime+" "+p.Prezime).Layout(gtx)
							})
						})
					}),
					layout.Flexed(0.6, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+w.text(projektant.Naziv)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+w.text(projektant.Ulica)+" "+w.text(projektant.Broj)+", "+w.text(projektant.Grad)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(projektant.JMBG)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(projektant.Email)).Layout(gtx)
							}))
					}))
			})
		},
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				id, _ := strconv.Atoi(klijentiIzbor.Value)
				klijent := w.Lica.Klijenti[id]
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.4, func(gtx C) D {
						return klijentiList.Layout(gtx, len(w.Lica.Klijenti), func(gtx C, i int) D {
							k := w.Lica.Klijenti[i]
							return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
								return material.RadioButton(w.UI.Tema.T, klijentiIzbor, strconv.Itoa(i), k.Naziv).Layout(gtx)
							})
						})
					}),
					layout.Flexed(0.6, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+w.text(klijent.Naziv)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+w.text(klijent.Ulica)+" "+w.text(klijent.Broj)+", "+w.text(klijent.Grad)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("PIB: ")+":"+w.text(klijent.PIB)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(klijent.MB)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(klijent.Email)).Layout(gtx)
							}))
					}))
			})
		},
	}
}
