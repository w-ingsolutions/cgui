package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

func (w *WingCal) GlavniEkran(gtx layout.Context) {
	if w.UI.Device == "p" {
		w.UI.TopSpace = 64
		w.UI.BottomSpace = 64
	} else {
		w.UI.TopSpace = 0
		w.UI.BottomSpace = 0
	}
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(helper.DuoUIline(false, 0, 0, w.UI.TopSpace, w.UI.Tema.Colors["Dark"])),
		layout.Flexed(1, func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(header(w)),
				layout.Flexed(1, w.strana(gtx)),
				layout.Rigid(footer(w)))
		}),
		layout.Rigid(helper.DuoUIline(false, 0, 0, w.UI.BottomSpace, w.UI.Tema.Colors["Dark"])))
}

func (w *WingCal) strana(gtx C) func(gtx C) D {
	switch d := gtx.Constraints.Max.X; {
	case d > 1910:
		w.UI.Device = "des"
	case d > 1356:
		w.UI.Device = "lap"
	case d > 758:
		w.UI.Device = "tab"
	case d < 758:
		w.UI.Device = "mob"
	}

	//case d > 2999:
	//w.UI.Device = "kkk"
	//case d > 1910:
	//w.UI.Device = "dtp"
	//case d > 1680:
	//w.UI.Device = "dtm"
	//case d > 1430:
	//w.UI.Device = "law"
	//case d > 1014:
	//w.UI.Device = "lap"
	//case d > 758:
	//w.UI.Device = "tbl"
	//case d > 558:
	//w.UI.Device = "tbp"
	//case d < 558:
	//w.UI.Device = "mob"
	//}
	var s func(gtx C) D
	prikazani := func(gtx C) D { return D{} }
	if w.Element {
		prikazani = w.PrikazaniElementSuma()
	}
	izbornikStrana := w.Panel(w.text("Radovi"), func(gtx C) D { return D{} }, w.IzbornikRadovaStrana(), prikazani)
	podesavanjaStrana := w.Panel(w.text("Podešavanja"), func(gtx C) D { return D{} }, w.PodesavanjaStrana(), w.sumaFooter(w.text("Podešavanja")))
	projekatStrana := w.Panel(w.text("Projekat"), func(gtx C) D { return D{} }, w.ProjekatStrana(), w.sumaFooter(w.text("Projekat")))
	materijalStrana := w.Panel(w.text("Materijal"), w.MaterijalStavke(), w.MaterijalStrana(), w.sumaFooter(w.text("Materijal")))
	sumaRadovaStrana := w.Panel(w.text("Ukupna cena radova"), w.SumaRadoviStavke(), w.SumaElementi(w.Suma.Elementi), w.sumaFooter(w.text("Suma: "+fmt.Sprintf("%.2f", w.Suma.SumaCena))))
	sumaMaterijalStrana := w.Panel(w.text("Ukupan neophodni materijal"), w.SumaStavkeMaterijala(), w.UkupanNeophodanMaterijal(w.Suma.NeophodanMaterijal), w.sumaFooter(w.text("Suma materijal: ")+fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal)))
	switch w.UI.Device {
	case "mob":
		switch w.Strana {
		case "materijal":
			s = materijalStrana
		case "projekat":
			s = projekatStrana
		case "radovi":
			s = izbornikStrana
		case "sumaRadovi":
			s = sumaRadovaStrana
		case "sumaMaterijal":
			s = sumaMaterijalStrana
		case "podesavanja":
			s = podesavanjaStrana
		}
	case "tab":
		//s = w.UI.SaMarginom.Layout(gtx, w.MaterijalStrana())
		s = w.Monitor(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana)
	case "lap":
		//s = w.UI.SaMarginom.Layout(gtx, w.MaterijalStrana())
		s = w.Monitor(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana)
	case "des":
		//s = w.UI.SaMarginom.Layout(gtx, w.MaterijalStrana())
		s = w.MonitorK(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana, materijalStrana, projekatStrana)
	}
	return s

	//return layout.Flex{
	//	Axis: layout.Horizontal,
	//}.Layout(w.UI.Context,
	//	layout.Flexed(1, func(gtx C) D {
	//		return w.UI.BezMargine.Layout(gtx, s)
	//	}))
}

func (w *WingCal) cell(align text.Alignment, tekst string) func(gtx C) D {
	return func(gtx C) D {
		return layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["LightGray"]).Layout(gtx, layout.N, func(gtx C) D {
				return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					cell := material.Caption(w.UI.Tema.T, tekst)
					cell.TextSize = unit.Dp(12)
					cell.Alignment = align
					return cell.Layout(gtx)
				})
			})
		})
	}
}

func (w *WingCal) sumaFooter(t string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		suma := material.Body2(w.UI.Tema.T, t)
		suma.Alignment = text.End
		return suma.Layout(gtx)
	}
}

func (w *WingCal) text(t string) string {
	return w.Kes.C(w.Jezik.t.T(latcyr.C(t, w.Podesavanja.Cyr)))
}
