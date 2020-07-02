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

			layout.Rigid(func(gtx C) D {
				return w.UI.Tema.WingUIcontainer(4, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.NW, func(gtx C) D {

					return layout.Center.Layout(gtx, func(gtx C) D {
						ic := w.UI.Tema.Icons["logo"]
						ic.Color = gelook.HexARGB("ffb8df42")
						return ic.Layout(gtx, unit.Dp(48))
					})
				})
			}),
			layout.Flexed(1, func(gtx C) D {

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
