package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
)

func header(w *WingCal) func(gtx C) D {
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
								//return layout.Center.Layout(gtx, func(gtx C) D {
								//	ic := w.UI.Tema.Icons["logo"]
								//	ic.Color = helper.HexARGB("ffb8df42")
								//	btn := material.IconButton(w.UI.Tema.T, projekatDugme, ic)
								//	btn.Color = helper.HexARGB(w.UI.Tema.Colors["ffb8df42"])
								//	btn.Size = unit.Dp(16)
								//	btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
								//	for projekatDugme.Clicked() {
								//		w.Strana = "Projekat"
								//	}
								//	return btn.Layout(gtx)
								//})
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
		})
	}
}
func (w *WingCal) headerMenu() []func(gtx C) D {
	return []func(gtx C) D{
		w.stranaDugme(radoviDugme, func() {}, w.text("RADOVI"), "radovi"),
		helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
		w.stranaDugme(sumaDugme, func() {}, w.text("SUMA"), "sumaRadovi"),
	}
}
