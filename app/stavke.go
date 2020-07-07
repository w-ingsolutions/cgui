package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
)

func (w *WingCal) SumaStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Flexed(0.5, w.cell(text.Start, w.text("Naziv"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Pojedinačna cena"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Količina"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.2, w.cell(text.End, w.text("Ukupna cena"))))
	}
}

func (w *WingCal) PrikazaniElementStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = width
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(0.4, w.cell(text.Start, w.text("Naziv"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Potrosnja"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Koeficijent"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Merodavna potrosnja"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.End, w.text("Cena materijala"))))
	}
}

func (w *WingCal) SumaRadoviStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.6, w.cell(text.Start, w.text("Naziv"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, w.text("Pojedinačna cena"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, w.text("Količina"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, w.text("Cena"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.05, w.cell(text.Middle, w.text("briši"))))
	}
}
func (w *WingCal) MaterijalStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.1, w.cell(text.Start, w.text("Id"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.3, w.cell(text.Middle, w.text("Naziv"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, w.text("Pakovanje"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, w.text("Jedinica mere"))),
			layout.Rigid(helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.2, w.cell(text.End, w.text("Cena"))))
	}
}
