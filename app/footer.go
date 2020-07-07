package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
)

func footer(w *WingCal) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
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
								btn := material.IconButton(w.UI.Tema.T, podesavanjeDugme, w.UI.Tema.Icons["settingsIcon"])
								btn.Color = helper.HexARGB(w.UI.Tema.Colors["Danger"])
								btn.Size = unit.Dp(16)
								btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
								for podesavanjeDugme.Clicked() {
									w.Strana = "podesavanja"
								}
								return btn.Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx C) D {
							return w.UI.BezMargine.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return headerMenuList.Layout(gtx, len(w.footerMenu()), func(gtx C, i int) D {
									return layout.UniformInset(unit.Dp(0)).Layout(gtx, w.footerMenu()[i])
								})
							})
						}))
				}))
		})
	}
}

func (w *WingCal) footerMenu() []func(gtx C) D {
	return []func(gtx C) D{
		w.stranaDugme(projekatDugme, w.text("PROJEKAT"), "projekat"),
		helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
		w.stranaDugme(sumaMaterialDugme, w.text("SUMA MATERIJAL"), "sumaMaterijal"),
		helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
		w.stranaDugme(materijalDugme, w.text("MATERIJAL"), "materijal"),
	}
}
