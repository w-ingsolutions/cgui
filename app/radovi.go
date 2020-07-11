package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
)

//func (w *WingCal) tIzborVrsteRadova() func(gtx C) D {
//		return w.vBtnLinks(
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[0]), w.btnLink(w.IzbornikRadova[1]), w.btnLink(w.IzbornikRadova[2])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[3]), w.btnLink(w.IzbornikRadova[4]), w.btnLink(w.IzbornikRadova[5])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[6]), w.btnLink(w.IzbornikRadova[7]), w.btnLink(w.IzbornikRadova[8])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[9]), w.btnLink(w.IzbornikRadova[10]), w.btnLink(w.IzbornikRadova[11])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[12]), w.btnLink(w.IzbornikRadova[13]), w.btnLink(w.IzbornikRadova[14])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[15]), w.btnLink(w.IzbornikRadova[16]), w.btnLink(w.IzbornikRadova[17])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[18]), w.btnLink(w.IzbornikRadova[19]), w.btnLink(w.IzbornikRadova[20])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[21]), w.btnLink(w.IzbornikRadova[22]), w.btnLink(w.IzbornikRadova[23])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[24]), w.btnLink(w.IzbornikRadova[25]), w.btnLink(w.IzbornikRadova[26])),
//			w.hBtnLinks(w.btnLink(w.IzbornikRadova[27]), w.btnLink(w.IzbornikRadova[28]), w.btnLink(w.IzbornikRadova[29])),
//		)
//}
//
//func (w *WingCal) thBtnLinks(a, b, c func(gtx C) D) func(gtx C) D {
//	return func(gtx C) D {
//		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
//			layout.Flexed(0.333, a),
//			layout.Flexed(0.333, b),
//			layout.Flexed(0.333, c))
//	}
//}

func (w *WingCal) IzborVrsteRadova() func(gtx C) D {
	return w.vBtnLinks(
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[0]), w.btnLink(w.IzbornikRadova[1])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[2]), w.btnLink(w.IzbornikRadova[3])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[4]), w.btnLink(w.IzbornikRadova[5])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[6]), w.btnLink(w.IzbornikRadova[7])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[8]), w.btnLink(w.IzbornikRadova[9])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[10]), w.btnLink(w.IzbornikRadova[11])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[12]), w.btnLink(w.IzbornikRadova[13])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[14]), w.btnLink(w.IzbornikRadova[15])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[16]), w.btnLink(w.IzbornikRadova[17])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[18]), w.btnLink(w.IzbornikRadova[19])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[20]), w.btnLink(w.IzbornikRadova[21])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[22]), w.btnLink(w.IzbornikRadova[23])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[24]), w.btnLink(w.IzbornikRadova[25])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[26]), w.btnLink(w.IzbornikRadova[27])),
		w.hBtnLinks(w.btnLink(w.IzbornikRadova[28]), w.btnLink(w.IzbornikRadova[29])),
	)
}

func (w *WingCal) hBtnLinks(a, b func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.333, a),
			layout.Flexed(0.333, b))
	}
}

func (w *WingCal) vBtnLinks(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {

		return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.1, a),
			layout.Flexed(0.1, b),
			layout.Flexed(0.1, c),
			layout.Flexed(0.1, d),
			layout.Flexed(0.1, e),
			layout.Flexed(0.1, f),
			layout.Flexed(0.1, g),
			layout.Flexed(0.1, h),
			layout.Flexed(0.1, i),
			layout.Flexed(0.1, j),
			layout.Flexed(0.1, k),
			layout.Flexed(0.1, l),
			layout.Flexed(0.1, m),
			layout.Flexed(0.1, n),
			layout.Flexed(0.1, o))
	}
}
func (w *WingCal) btnLink(r model.ElementMenu) func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, r.Link, fmt.Sprint(r.Id)+". "+w.text(r.Title))
		btn.CornerRadius = unit.Dp(0)
		btn.Inset = layout.Inset{
			Top:    unit.Dp(2),
			Right:  unit.Dp(2),
			Bottom: unit.Dp(2),
			Left:   unit.Dp(2),
		}
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Gray"])
		if r.Materijal {
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["DarkGray"])
		}
		w.LinkoviIzboraVrsteRadovaKlik(r)
		return btn.Layout(gtx)
	}
}
