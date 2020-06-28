package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/gelook"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

var (
	latcyrDugme    = new(widget.Clickable)
	izbornikDugme  = new(widget.Clickable)
	materijalDugme = new(widget.Clickable)
	sumaDugme      = new(widget.Clickable)

	putanjaRadoviDugme   = new(widget.Clickable)
	putanjaPodvrstaDugme = new(widget.Clickable)
	putanjaElementDugme  = new(widget.Clickable)

	headerMenuList = &layout.List{
		Axis: layout.Horizontal,
	}
)

func header(w *WingCal) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(0.5, func(gtx C) D {
				return putanjaList.Layout(gtx, len(w.Putanja), func(gtx C, i int) D {
					var d D
					//switch i {
					//case 0:
					//	return	material.Button(w.UI.Tema.T, putanjaRadoviDugme, latcyr.C(w.Putanja[i], w.Cyr)).Layout(gtx)
					//	for putanjaRadoviDugme.Clicked() {
					//		komanda := ""
					//		//if len(w.Putanja) == 3 {
					//		//	komanda = "/" + fmt.Sprint(w.Roditelj)
					//		//	//podvrstaradova = fmt.Sprint(w.Roditelj)
					//		//	fmt.Println("roddddditeL111::", w.Roditelj)
					//		//}
					//		//if len(w.Putanja) == 4 {
					//		//	komanda = "/" + calc.Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
					//		//}
					//		w.APIpozivIzbornik("radovi" + komanda)
					//		//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
					//		w.GenerisanjeLinkova(w.IzbornikRadova)
					//		//w.Putanja = nil
					//		//w.Putanja = append(w.Putanja, w.Radovi.Naziv)
					//
					//	}
					//case 1:
					//	return material.Button(w.UI.Tema.T, putanjaPodvrstaDugme, latcyr.C(w.Putanja[i], w.Cyr)).Layout(gtx)
					//	for putanjaPodvrstaDugme.Clicked() {
					//		komanda := ""
					//
					//		komanda = "/" + fmt.Sprint(w.Roditelj)
					//		//podvrstaradova = fmt.Sprint(w.Roditelj)
					//		//fmt.Println("roddddditeL111::", w.Roditelj)
					//		//komanda = "/" + calc.Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
					//
					//		w.APIpozivIzbornik("radovi" + komanda)
					//		//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
					//		w.GenerisanjeLinkova(w.IzbornikRadova)
					//		w.Putanja = w.Putanja[:len(w.Putanja)-1]
					//
					//	}
					//case 2:
					//	return material.Button(w.UI.Tema.T, putanjaElementDugme,latcyr.C(w.Putanja[i], w.Cyr)).Layout(gtx)
					//	for putanjaElementDugme.Clicked() {
					//		w.Element = false
					//	}
					//}
					return d
				})
			}),
			layout.Flexed(0.5, func(gtx C) D {

				headerMenu := []func(gtx C) D{
					func(gtx C) D {
						btnKalkulator := material.Button(w.UI.Tema.T, izbornikDugme, latcyr.C("IZBORNIK", w.Cyr))
						btnKalkulator.Background = gelook.HexARGB(w.UI.Tema.Colors["Secondary"])
						for izbornikDugme.Clicked() {
							w.Strana = "izbornik"
						}
						return btnKalkulator.Layout(gtx)
					},

					func(gtx C) D {
						btnMaterijal := material.Button(w.UI.Tema.T, sumaDugme, latcyr.C("SUMA", w.Cyr))
						btnMaterijal.Background = gelook.HexARGB(w.UI.Tema.Colors["Secondary"])
						for sumaDugme.Clicked() {
							w.Strana = "suma"
						}
						return btnMaterijal.Layout(gtx)
					},
					func(gtx C) D {
						btnMaterijal := material.Button(w.UI.Tema.T, materijalDugme, latcyr.C("MATERIJAL", w.Cyr))
						btnMaterijal.Background = gelook.HexARGB(w.UI.Tema.Colors["Secondary"])
						for materijalDugme.Clicked() {
							w.Strana = "materijal"
						}
						return btnMaterijal.Layout(gtx)
					},
				}
				return headerMenuList.Layout(gtx, len(headerMenu), func(gtx C, i int) D {
					return layout.UniformInset(unit.Dp(0)).Layout(gtx, headerMenu[i])
				})
			}),
			layout.Rigid(func(gtx C) D {

				latcyr := "Ћирилица"
				if w.Cyr {
					latcyr = "Latinica"
				}

				btnLatcyr := material.Button(w.UI.Tema.T, latcyrDugme, latcyr)
				btnLatcyr.Background = gelook.HexARGB(w.UI.Tema.Colors["Secondary"])
				for latcyrDugme.Clicked() {
					if w.Cyr {
						w.Cyr = false
					} else {
						w.Cyr = true
					}
				}
				return btnLatcyr.Layout(gtx)
			}))
	}
}
