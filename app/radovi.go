package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/icontextbtn"
	"github.com/w-ingsolutions/c/model"
)

/////////////
func (w *WingCal) IzborVrsteRadova() func(gtx C) D {
	izbornik := w.IzborPodVrsteRadova()
	switch w.UI.Device {
	case "mob":
		izbornik = w.IzborPodVrsteRadova()
	case "tab":
		izbornik = w.IzborVrsteRadovaDveKolone()
	case "lap":
		izbornik = w.IzborVrsteRadovaTriKolone()
	case "des":
		izbornik = w.IzborVrsteRadovaCetiriKolone()
	}
	return izbornik
}
func (w *WingCal) IzborVrsteRadovaDveKolone() func(gtx C) D {
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
			layout.Flexed(0.06, a),
			layout.Flexed(0.06, b),
			layout.Flexed(0.06, c),
			layout.Flexed(0.06, d),
			layout.Flexed(0.06, e),
			layout.Flexed(0.06, f),
			layout.Flexed(0.06, g),
			layout.Flexed(0.06, h),
			layout.Flexed(0.06, i),
			layout.Flexed(0.06, j),
			layout.Flexed(0.06, k),
			layout.Flexed(0.06, l),
			layout.Flexed(0.06, m),
			layout.Flexed(0.06, n),
			layout.Flexed(0.06, o))

	}
}

func (w *WingCal) IzborVrsteRadovaTriKolone() func(gtx C) D {
	return w.tvBtnLinks(
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[0]), w.btnLink(w.IzbornikRadova[1]), w.btnLink(w.IzbornikRadova[2])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[3]), w.btnLink(w.IzbornikRadova[4]), w.btnLink(w.IzbornikRadova[5])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[6]), w.btnLink(w.IzbornikRadova[7]), w.btnLink(w.IzbornikRadova[8])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[9]), w.btnLink(w.IzbornikRadova[10]), w.btnLink(w.IzbornikRadova[11])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[12]), w.btnLink(w.IzbornikRadova[13]), w.btnLink(w.IzbornikRadova[14])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[15]), w.btnLink(w.IzbornikRadova[16]), w.btnLink(w.IzbornikRadova[17])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[18]), w.btnLink(w.IzbornikRadova[19]), w.btnLink(w.IzbornikRadova[20])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[21]), w.btnLink(w.IzbornikRadova[22]), w.btnLink(w.IzbornikRadova[23])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[24]), w.btnLink(w.IzbornikRadova[25]), w.btnLink(w.IzbornikRadova[26])),
		w.thBtnLinks(w.btnLink(w.IzbornikRadova[27]), w.btnLink(w.IzbornikRadova[28]), w.btnLink(w.IzbornikRadova[29])),
	)
}

func (w *WingCal) thBtnLinks(a, b, c func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.333, a),
			layout.Flexed(0.333, b),
			layout.Flexed(0.333, c))
	}
}

func (w *WingCal) tvBtnLinks(a, b, c, d, e, f, g, h, i, j func(gtx C) D) func(gtx C) D {
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
			layout.Flexed(0.1, j))
	}
}

func (w *WingCal) IzborVrsteRadovaCetiriKolone() func(gtx C) D {
	return w.fvBtnLinks(
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[0]), w.btnLink(w.IzbornikRadova[1]), w.btnLink(w.IzbornikRadova[2]), w.btnLink(w.IzbornikRadova[3])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[4]), w.btnLink(w.IzbornikRadova[5]), w.btnLink(w.IzbornikRadova[6]), w.btnLink(w.IzbornikRadova[7])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[8]), w.btnLink(w.IzbornikRadova[9]), w.btnLink(w.IzbornikRadova[10]), w.btnLink(w.IzbornikRadova[11])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[12]), w.btnLink(w.IzbornikRadova[13]), w.btnLink(w.IzbornikRadova[14]), w.btnLink(w.IzbornikRadova[15])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[16]), w.btnLink(w.IzbornikRadova[17]), w.btnLink(w.IzbornikRadova[18]), w.btnLink(w.IzbornikRadova[19])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[20]), w.btnLink(w.IzbornikRadova[21]), w.btnLink(w.IzbornikRadova[22]), w.btnLink(w.IzbornikRadova[23])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[24]), w.btnLink(w.IzbornikRadova[25]), w.btnLink(w.IzbornikRadova[26]), w.btnLink(w.IzbornikRadova[27])),
		w.fhBtnLinks(w.btnLink(w.IzbornikRadova[28]), w.btnLink(w.IzbornikRadova[29]), w.btnLink(w.IzbornikRadova[28]), w.btnLink(w.IzbornikRadova[29])),
	)
}

func (w *WingCal) fhBtnLinks(a, b, c, d func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.25, a),
			layout.Flexed(0.25, b),
			layout.Flexed(0.25, c),
			layout.Flexed(0.25, d))
	}
}

func (w *WingCal) fvBtnLinks(a, b, c, d, e, f, g, h func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.125, a),
			layout.Flexed(0.125, b),
			layout.Flexed(0.125, c),
			layout.Flexed(0.125, d),
			layout.Flexed(0.125, e),
			layout.Flexed(0.125, f),
			layout.Flexed(0.125, g),
			layout.Flexed(0.125, h))

	}
}

func (w *WingCal) btnLink(r model.ElementMenu) func(gtx C) D {
	return func(gtx C) D {
		btn := icontextbtn.IconTextBtn(w.UI.Tema.T, r.Link, w.UI.Tema.Icons[r.Slug], unit.Dp(48), w.UI.Tema.Colors["Light"], fmt.Sprint(r.Id)+". "+w.text(r.Title))
		switch w.UI.Device {
		case "mob":
			btn.Axis = layout.Horizontal
		case "tab":
			btn.Axis = layout.Horizontal
		case "lap":
			btn.IconSize = unit.Dp(64)
			btn.TextSize = unit.Dp(12)
			btn.Axis = layout.Vertical

		case "des":
			btn.IconSize = unit.Dp(72)
			btn.TextSize = unit.Dp(12)
			btn.Axis = layout.Vertical
		}
		btn.CornerRadius = unit.Dp(0)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Gray"])
		//btn.Inset = layout.Inset{
		//	Top:    unit.Dp(2),
		//	Right:  unit.Dp(2),
		//	Bottom: unit.Dp(2),
		//	Left:   unit.Dp(2),
		//}
		if r.Materijal {
			//btn.Background = helper.HexARGB(w.UI.Tema.Colors["DarkGray"])
		}
		w.LinkoviIzboraVrsteRadovaKlik(r)
		return btn.Layout(gtx)
	}
}
