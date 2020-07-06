package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
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
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   layout.SpaceBetween,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							ic := w.UI.Tema.Icons["logo"]
							ic.Color = helper.HexARGB("ffb8df42")
							return ic.Layout(gtx, unit.Dp(32))
						})
					}),
					layout.Rigid(func(gtx C) D {
						return w.UI.BezMargine.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return headerMenuList.Layout(gtx, len(w.headerMenu()), func(gtx C, i int) D {
								return layout.UniformInset(unit.Dp(0)).Layout(gtx, w.headerMenu()[i])
							})
						})
					}))
			}))
	}
}

func (w *WingCal) headerMenu() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			btn := material.Button(w.UI.Tema.T, izbornikDugme, latcyr.C("RADOVI", w.Cyr))
			btn.CornerRadius = unit.Dp(0)
			btn.TextSize = unit.Dp(12)
			btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
			for izbornikDugme.Clicked() {
				w.Strana = "radovi"
			}
			return btn.Layout(gtx)
		},
		helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
		func(gtx C) D {
			btn := material.Button(w.UI.Tema.T, sumaDugme, latcyr.C("SUMA", w.Cyr))
			btn.CornerRadius = unit.Dp(0)
			btn.TextSize = unit.Dp(12)
			btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
			for sumaDugme.Clicked() {
				w.Strana = "suma"
			}
			return btn.Layout(gtx)
		},
		helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
		func(gtx C) D {
			btn := material.Button(w.UI.Tema.T, materijalDugme, latcyr.C("MATERIJAL", w.Cyr))
			btn.CornerRadius = unit.Dp(0)
			btn.TextSize = unit.Dp(12)
			btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
			for materijalDugme.Clicked() {
				w.Strana = "materijal"
			}
			return btn.Layout(gtx)
		},
	}
}
