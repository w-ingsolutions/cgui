package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/counter"
	"github.com/w-ingsolutions/c/pkg/latcyr"
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
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Projektant")).Layout(gtx) },
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				id, _ := strconv.Atoi(projektantIzbor.Value)
				projektant := w.Lica.Projektanti[id]
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.4, func(gtx C) D {
						return projektantiList.Layout(gtx, len(w.Lica.Projektanti), func(gtx C, i int) D {
							p := w.Lica.Projektanti[i]
							return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
								return material.RadioButton(w.UI.Tema.T, projektantIzbor, strconv.Itoa(i), p.KratakNaziv).Layout(gtx)
							})
						})
					}),
					layout.Flexed(0.6, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+w.text(projektant.KratakNaziv)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+w.text(projektant.Adresa)+" "+w.text(projektant.Grad)).Layout(gtx)
								//}),
								//layout.Rigid(func(gtx C) D {
								//return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(projektant.JMBG)).Layout(gtx)
								//}),
								//layout.Rigid(func(gtx C) D {
								//	return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(projektant.Email)).Layout(gtx)
							}))
					}))
			})
		},
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Investitor")).Layout(gtx) },
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				id, _ := strconv.Atoi(klijentiIzbor.Value)
				investotor := w.Lica.Investotori[id]
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.4, func(gtx C) D {
						return klijentiList.Layout(gtx, len(w.Lica.Investotori), func(gtx C, i int) D {
							inv := w.Lica.Investotori[i]
							return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
								return material.RadioButton(w.UI.Tema.T, klijentiIzbor, strconv.Itoa(i), inv.KratakNaziv).Layout(gtx)
							})
						})
					}),
					layout.Flexed(0.6, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+w.text(investotor.KratakNaziv)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+w.text(investotor.Adresa)+" "+w.text(investotor.Grad)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("PIB: ")+":"+w.text(investotor.PIB)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(investotor.MB)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Datum osnivanja: ")+":"+w.text(investotor.DatumOsnivanja)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Delatnost: ")+":"+w.text(investotor.Delatnost)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Računi: ")).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return racuniList.Layout(gtx, len(investotor.Racuni), func(gtx C, i int) D {
									racuni := investotor.Racuni[i]
									return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
										return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
											layout.Rigid(func(gtx C) D {
												return material.Caption(w.UI.Tema.T, w.text(racuni.Banka)).Layout(gtx)
											}),
											layout.Rigid(func(gtx C) D {
												return material.Caption(w.UI.Tema.T, w.text(racuni.Racun)).Layout(gtx)
											}))
									})
								})
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(investotor.Email)).Layout(gtx)
							}))
					}))
			})
		},
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Radovi")).Layout(gtx) },
		func(gtx C) D {
			projekat.CenaRadova = w.Suma.SumaCena + w.Suma.SumaCena*float64(radovi.Value)/100
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.5, func(gtx C) D {
						return material.H6(w.UI.Tema.T, w.text("Ukupna cena radova")).Layout(gtx)
					}),
					layout.Flexed(0.5, func(gtx C) D {
						return material.H6(w.UI.Tema.T, fmt.Sprintf("%.2f", projekat.CenaRadova)).Layout(gtx)
					}),
					layout.Rigid(counter.DuoUIcounterSt(w.UI.Tema, radovi, func() {}).Layout(radovi, gtx, w.UI.Tema.T, latcyr.C("%", w.Podesavanja.Cyr), fmt.Sprint(radovi.Value))))
			})
		},
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Materijal")).Layout(gtx) },
		func(gtx C) D {
			projekat.CenaMaterijala = w.Suma.SumaCenaMaterijal + w.Suma.SumaCenaMaterijal*float64(materijal.Value)/100
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
				return layout.Flex{}.Layout(gtx,
					layout.Flexed(0.5, func(gtx C) D {
						return material.H6(w.UI.Tema.T, w.text("Ukupna cena materijala")).Layout(gtx)
					}),
					layout.Flexed(0.5, func(gtx C) D {
						return material.H6(w.UI.Tema.T, fmt.Sprintf("%.2f", projekat.CenaMaterijala)).Layout(gtx)
					}),
					layout.Rigid(counter.DuoUIcounterSt(w.UI.Tema, materijal, func() {}).Layout(materijal, gtx, w.UI.Tema.T, latcyr.C("%", w.Podesavanja.Cyr), fmt.Sprint(materijal.Value))))
			})
		},
	}
}
