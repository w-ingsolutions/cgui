package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
)

func (w *WingCal) RadNeophodanMaterijal(l *layout.List) func(gtx C) D {
	return func(gtx C) D {
		//var materijal model.WingNeophodanMaterijal
		nm := w.PrikazaniElement.NeophodanMaterijal
		//width := gtx.Constraints.Max.X
		return l.Layout(gtx, len(nm), func(gtx C, i int) D {
			materijal := nm[i]
			id := materijal.Id - 1
			materijal.Koeficijent = materijal.Koeficijent
			materijal.Materijal = *w.Materijal[id]
			if materijal.Koeficijent > 0 {
				materijal.Kolicina = materijal.Materijal.Potrosnja * float64(w.UI.Counters.Kolicina.Value) * materijal.Koeficijent
			}
			materijal.UkupnaCena = materijal.Materijal.Cena * float64(materijal.Kolicina)
			materijal.UkupnoPakovanja = int(materijal.Kolicina / float64(materijal.Materijal.Pakovanje))

			//gtx.Constraints.Min.X = width
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Spacing:   layout.SpaceBetween,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Flexed(0.4, w.cell(text.Start, w.text(materijal.Materijal.Naziv))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Materijal.Potrosnja))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.Middle, fmt.Sprint(materijal.Koeficijent))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Kolicina))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.End, fmt.Sprintf("%.2f", materijal.UkupnaCena))),
					)
				}),
				layout.Rigid(helper.DuoUIline(false, 1, 1, 1, w.UI.Tema.Colors["Gray"])))
		})
	}
}
func (w *WingCal) UkupanNeophodanMaterijal(unm map[int]model.WingNeophodanMaterijal) func(gtx C) D {
	return func(gtx C) D {
		width := gtx.Constraints.Max.X
		return ukupanNeophodanMaterijalList.Layout(gtx, len(unm), func(gtx C, i int) D {
			//materijal := unm[i]
			materijal := w.Suma.NeophodanMaterijalPrikaz[i]
			gtx.Constraints.Min.X = width
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Spacing:   layout.SpaceBetween,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Flexed(0.5, w.cell(text.Start, w.text(materijal.Materijal.Naziv))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.Middle, fmt.Sprint(materijal.Materijal.Cena))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.15, w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Kolicina))),
						layout.Rigid(helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"])),
						layout.Flexed(0.2, w.cell(text.End, fmt.Sprintf("%.2f", materijal.UkupnaCena))))
				}),
				layout.Rigid(helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"])),
			)
		})
	}
}
