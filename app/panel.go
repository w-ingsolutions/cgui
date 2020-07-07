package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
)

func (w *WingCal) Panel(naslov string, stavke, sadrzaj, footer func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return w.UI.SaMarginom.Layout(gtx, func(gtx C) D {
			return w.UI.SaMarginom.Layout(gtx, func(gtx C) D {
				return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
					return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Flexed(1, func(gtx C) D {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
											gtx.Constraints.Min.X = gtx.Constraints.Max.X
											suma := material.H6(w.UI.Tema.T, naslov)
											suma.Alignment = text.Start
											return suma.Layout(gtx)
										})

									}),
									layout.Rigid(func(gtx C) D {
										return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
											layout.Rigid(func(gtx C) D {
												return layout.UniformInset(unit.Dp(4)).Layout(gtx, stavke)
											}))
									}),
									//layout.Rigid(helper.DuoUIline(false, 4, 4, 2, w.UI.Tema.Colors["Primary"])),
									layout.Flexed(1, sadrzaj))
							}),
							layout.Rigid(func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, footer)
							}),
						)
					})
				})
			})
		})
	}
}
