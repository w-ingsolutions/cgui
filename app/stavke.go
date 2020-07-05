package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) SumaStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Flexed(0.5, w.cell(text.Start, latcyr.C("Naziv", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Pojedinačna cena", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Količina", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.2, w.cell(text.End, latcyr.C("Ukupna cena", w.Cyr))))
	}
}

func (w *WingCal) PrikazaniElementStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = width
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Flexed(0.4, w.cell(text.Start, latcyr.C("Naziv", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Potrosnja", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Koeficijent", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Merodavna potrosnja", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.End, latcyr.C("Cena materijala", w.Cyr))))
	}
}

func (w *WingCal) SumaRadoviStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.6, w.cell(text.Start, latcyr.C("Naziv", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.15, w.cell(text.Middle, latcyr.C("Pojedinačna cena", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, latcyr.C("Količina", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, latcyr.C("Cena", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.05, w.cell(text.Middle, latcyr.C("briši", w.Cyr))))
	}
}
func (w *WingCal) MaterijalStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(0.1, w.cell(text.Start, latcyr.C("Id", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.3, w.cell(text.Middle, latcyr.C("Naziv", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, latcyr.C("Pakovanje", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.1, w.cell(text.Middle, latcyr.C("Jedinica mere", w.Cyr))),
			layout.Rigid(w.UI.Tema.WingUIline(true, 0, 0, 2, w.UI.Tema.Colors["Gray"])),
			layout.Flexed(0.2, w.cell(text.End, latcyr.C("Cena", w.Cyr))))
	}
}
