package calc

import (
	"fmt"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/skratchdot/open-golang/open"
	"github.com/w-ingsolutions/c/pkg/pdf"
)

func (w *WingCal) Stampa() func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, stampajDugme, w.text("Štampaj"))
		btn.CornerRadius = unit.Dp(0)

		if len(w.Suma.Elementi) != 0 {
			for stampajDugme.Clicked() {

				p := pdf.P()

				tr := p.UnicodeTranslatorFromDescriptor("")

				p.SetTopMargin(30)

				marginCell := 2. // margin of top/bottom of cell
				pagew, pageh := p.GetPageSize()
				mleft, mright, _, mbottom := p.GetMargins()

				p.SetHeaderFuncMode(func() {
					//pdf.Image("/usr/home/marcetin/Public/wingcal/NOVOGUI/pdfheader.png", 5, 5, 200, 25, false, "", 0, "")
					//pdf.SetDrawColor(200,200,200)
					p.SetFillColor(200, 200, 200)
					p.Rect(5, 5, 200, 20, "F")
					p.SetY(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "MB:20701005", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "B", 10)
					p.CellFormat(47, 10, "W-ing Solutions D.o.o.", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     SIFRA PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "PIB:106892584", "0", 0, "L", false, 0, "")
					p.CellFormat(47, 8, "Bulevar Oslobodjenja 30A", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     NAZIV PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "projekat za evidentiranje", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
					p.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     DATUM PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 10, "mart 2020", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")

				}, true)
				p.SetFooterFunc(func() {

					p.SetY(-15)
					p.SetFont("Arial", "I", 8)
					p.CellFormat(0, 10, fmt.Sprintf("Strana %d/{nb}", p.PageNo()),
						"", 0, "C", false, 0, "")
					p.Ln(0)
					p.SetFillColor(200, 200, 200)
					p.Rect(5, 275, 200, 20, "F")
					p.SetY(-22)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "MB:20701005", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "B", 10)
					p.CellFormat(47, 10, "W-ing Solutions D.o.o.", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     SIFRA PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "PIB:106892584", "0", 0, "L", false, 0, "")
					p.CellFormat(47, 8, "Bulevar Oslobodjenja 30A", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     NAZIV PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "projekat za evidentiranje", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
					p.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 8, "     DATUM PROJEKTA", "0", 0, "L", false, 0, "")
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 10, "mart 2020", "0", 0, "R", false, 0, "")
					p.Ln(5)
					p.SetFont("Arial", "", 8)
					p.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")

				})
				p.AliasNbPages("")

				p.AddPage()
				p.SetFont("Times", "B", 16)
				p.CellFormat(0, 10, tr(w.text("Specifikacija aktivnosti")), "0", 0, "", false, 0, "")
				p.Ln(20)

				p.SetFont("Arial", "", 10)
				for _, e := range w.Suma.Elementi {
					cols := []float64{40, pagew - mleft - mright - 20}
					//rows := [][]string{}

					rows := [][]string{
						[]string{
							"Šifra", e.Sifra,
						},
						[]string{
							"Naziv", e.Element.Naziv,
						},
						[]string{
							"Opis", e.Element.Opis,
						},
						[]string{
							"Jedinica mere", e.Element.Jedinica,
						},
						[]string{
							"Jedinična cena", fmt.Sprint(e.Element.Cena),
						},
						[]string{
							"Količina", fmt.Sprint(e.Kolicina),
						},
						[]string{
							"Vrednost rada", fmt.Sprintf("%.2f", e.SumaCena),
						},
					}
					for _, row := range rows {
						curx, y := p.GetXY()
						x := curx
						height := 0.
						_, lineHt := p.GetFontSize()
						for i, txt := range row {
							lines := p.SplitLines([]byte(txt), cols[i])
							h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
							if h > height {
								height = h
							}
						}
						// add a new page if the height of the row doesn't fit on the page
						if p.GetY()+height > pageh-mbottom {
							p.AddPage()
							y = p.GetY()
						}
						for i, txt := range row {
							width := cols[i]
							//pdf.Rect(x, y, width, height, "")
							if i < 1 {
								p.SetFont("Arial", "B", 10)
							} else {
								p.SetFont("Arial", "", 10)
							}
							//fmt.Println("Col::", i)

							p.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
							x += width
							p.SetXY(x, y)
						}
						p.SetXY(curx, y+height)
					}
					p.Ln(8)
				}

				p.SetFont("Times", "B", 16)
				p.CellFormat(0, 10, w.text("Suma: ")+fmt.Sprintf("%.2f", w.Suma.SumaCena), "0", 0, "", false, 0, "")
				p.Ln(40)
				p.AddPage()

				p.SetFont("Times", "B", 16)
				p.CellFormat(0, 10, w.text("Specifikacija materijala"), "0", 0, "", false, 0, "")
				p.Ln(20)

				p.SetFont("Arial", "", 10)
				for _, e := range w.Suma.UkupanNeophodanMaterijalPrikaz {
					cols := []float64{40, pagew - mleft - mright - 20}
					//rows := [][]string{}

					rows := [][]string{
						[]string{
							"Šifra", fmt.Sprint(e.Id),
						},
						[]string{
							"Naziv", e.Materijal.Naziv,
						},
						//[]string{
						//	"Osobine i namena", e.Materijal.OsobineNamena,
						//},
						[]string{
							"Jedinica mere", e.Materijal.JedinicaPotrosnje,
						},
						[]string{
							"Jedinična cena", fmt.Sprint(e.Materijal.Cena),
						},
						[]string{
							"Količina", fmt.Sprint(e.Kolicina),
						},
						[]string{
							"Vrednost materijala", fmt.Sprintf("%.2f", e.UkupnaCena),
						},
					}
					for _, row := range rows {
						curx, y := p.GetXY()
						x := curx
						height := 0.
						_, lineHt := p.GetFontSize()
						for i, txt := range row {
							lines := p.SplitLines([]byte(txt), cols[i])
							h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
							if h > height {
								height = h
							}
						}
						// add a new page if the height of the row doesn't fit on the page
						if p.GetY()+height > pageh-mbottom {
							p.AddPage()
							y = p.GetY()
						}
						for i, txt := range row {
							width := cols[i]
							//pdf.Rect(x, y, width, height, "")
							p.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
							x += width
							p.SetXY(x, y)
						}
						p.SetXY(curx, y+height)
					}
					p.Ln(8)
				}

				p.SetFont("Times", "B", 16)
				p.CellFormat(0, 10, w.text("Suma materijal: ")+fmt.Sprintf("%.2f", projekat.CenaMaterijala), "0", 0, "", false, 0, "")
				p.Ln(20)

				///////////////////////////////////

				p.AddPage()

				p.SetFont("Times", "B", 16)
				p.CellFormat(0, 10, w.text("Tehnički list"), "0", 0, "", false, 0, "")
				p.Ln(20)

				p.SetFont("Arial", "", 10)
				for _, e := range w.Suma.UkupanNeophodanMaterijalPrikaz {
					cols := []float64{40, pagew - mleft - mright - 20}
					//rows := [][]string{}

					rows := [][]string{
						[]string{
							"Šifra", fmt.Sprint(e.Id),
						},
						[]string{
							"Naziv", e.Materijal.Naziv,
						},
						[]string{
							"Osobine i namena", e.Materijal.OsobineNamena,
						},
						[]string{
							"Nacin rada", e.Materijal.NacinRada,
						},

						[]string{
							"Jedinica mere", e.Materijal.JedinicaPotrosnje,
						},
					}
					for _, row := range rows {
						curx, y := p.GetXY()
						x := curx
						height := 0.
						_, lineHt := p.GetFontSize()
						for i, txt := range row {
							lines := p.SplitLines([]byte(txt), cols[i])
							h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
							if h > height {
								height = h
							}
						}
						// add a new page if the height of the row doesn't fit on the page
						if p.GetY()+height > pageh-mbottom {
							p.AddPage()
							y = p.GetY()
						}
						for i, txt := range row {
							width := cols[i]
							//pdf.Rect(x, y, width, height, "")
							p.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
							x += width
							p.SetXY(x, y)
						}
						p.SetXY(curx, y+height)
					}
					p.Ln(8)
				}

				//////////////////

				//// See documentation for details on how to generate fonts
				////pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
				//
				//fontSize := 16.0
				////pdf.SetFont("Helvetica", "", fontSize)
				//ht := pdf.PointConvert(fontSize)
				//tr := pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
				////write := func(str string) {
				////	pdf.CellFormat(190, ht, tr(str), "", 1, "C", false, 0, "")
				//	//pdf.MultiCell(190, ht, tr(str), "", "C", false)
				//	//pdf.Ln(ht)
				////}
				//pdf.AddPage()
				//str := ""
				//
				//pdf.MultiCell(190, ht, str, "", "L", false)
				//pdf.Ln(2 * ht)
				//
				////write("Srpski testiranje latinicei tako čćž qww yxx šđ šđ žć")
				//
				////pdf.SetFont("Helvetica-1251", "", fontSize) // Name matches one specified in AddFont()
				////tr = pdf.UnicodeTranslatorFromDescriptor("cp1251")
				////write("Тестирање на Српски од че до ће и све ћж љљљ жџџџџ и тако те")
				//
				//pdf.MultiCell(190, ht, tr("Тестирање на Српски од че до ће и све ћж љљљ жџџџџ и тако те"), "", "C", false)
				//pdf.Ln(ht)
				//
				//pdf.MultiCell(190, ht, tr("Srpski testiranje latinicei tako čćž qww yxx šđ šđ žć"), "", "C", false)
				//pdf.Ln(ht)

				// Output:
				// Successfully generated pdf/Fpdf_CellFormat_codepage.pdf

				err := p.OutputFileAndClose(w.Podesavanja.Dir + "/nalog.pdf")
				if err != nil {
				}
				open.Run(w.Podesavanja.Dir + "/nalog.pdf")

			}
		}
		return btn.Layout(gtx)
	}
}
