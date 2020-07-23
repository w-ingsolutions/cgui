package calc

import (
	"fmt"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/jung-kurt/gofpdf"
	"github.com/skratchdot/open-golang/open"
	"github.com/w-ingsolutions/c/pkg/pdf"
	"time"
)

var (
	materijal, aktivnosti string
)

func (w *WingCal) Stampa() func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, stampajDugme, w.text("Štampaj"))
		btn.CornerRadius = unit.Dp(0)
		if len(w.Suma.Elementi) != 0 {
			for stampajDugme.Clicked() {
				fmt.Println("proj::", projekat.Investitor)
				p := pdf.P()

				tr := p.UnicodeTranslatorFromDescriptor("")

				p.SetTopMargin(30)
				marginCell := 2. // margin of top/bottom of cell
				pagew, pageh := p.GetPageSize()
				mleft, mright, _, mbottom := p.GetMargins()

				p.SetHeaderFuncMode(w.pdfHeader(p), true)
				p.SetFooterFunc(w.pdfFooter(p))
				p.AliasNbPages("")
				w.ponuda(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
				w.ipList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)

				w.specifikacijaRadovaList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
				w.specifikacijaMaterijalaList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
				w.tehnickiList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)

				//err := p.OutputFileAndClose(w.Podesavanja.Dir + "/nalog.pdf")
				err := p.OutputFileAndClose("nalog.pdf")
				if err != nil {
				}
				//open.Run(w.Podesavanja.Dir + "/nalog.pdf")
				open.Run("nalog.pdf")

			}
		}
		return btn.Layout(gtx)
	}
}

func (w *WingCal) pdfHeader(p *gofpdf.Fpdf) func() {
	return func() {
		currentTime := time.Now()
		//pdf.Image("/usr/home/marcetin/Public/wingcal/NOVOGUI/pdfheader.png", 5, 5, 200, 25, false, "", 0, "")
		//pdf.SetDrawColor(200,200,200)
		p.SetFillColor(200, 200, 200)
		p.Rect(5, 5, 200, 20, "F")
		p.SetY(5)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "MB:20701005", "0", 0, "L", false, 0, "")
		p.SetFont("Arial", "B", 10)
		p.CellFormat(47, 10, projekat.Projektant.KratakNaziv, "0", 0, "R", false, 0, "")
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
		p.CellFormat(47, 10, currentTime.Format("06-Jan-02"), "0", 0, "R", false, 0, "")
		p.Ln(5)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")
	}
}
func (w *WingCal) pdfFooter(p *gofpdf.Fpdf) func() {
	return func() {
		p.SetY(-15)
		p.SetFont("Arial", "I", 8)
		p.CellFormat(0, 10, fmt.Sprintf("Strana %d/{nb}", p.PageNo()),
			"", 0, "C", false, 0, "")
		p.Ln(0)
		p.SetFillColor(200, 200, 200)
		p.Rect(5, 275, 200, 20, "F")
		p.SetY(-22)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "MB:"+projekat.Investitor.MB, "0", 0, "L", false, 0, "")
		p.SetFont("Arial", "B", 10)
		p.CellFormat(47, 10, projekat.Investitor.KratakNaziv, "0", 0, "R", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 8, "     SIFRA DOKUMENTA", "0", 0, "L", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
		p.Ln(5)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "PIB:"+projekat.Investitor.PIB, "0", 0, "L", false, 0, "")
		p.CellFormat(47, 8, projekat.Investitor.Adresa, "0", 0, "R", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 8, "     NAZIV DOKUMENTA", "0", 0, "L", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 8, "dokument za evidentiranje", "0", 0, "R", false, 0, "")
		p.Ln(5)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
		p.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 8, "     DATUM DOKUMENTA", "0", 0, "L", false, 0, "")
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 10, "mart 2020", "0", 0, "R", false, 0, "")
		p.Ln(5)
		p.SetFont("Arial", "", 8)
		p.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")

	}
}

func (w *WingCal) tehnickiList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Tehnički list"), "0", 0, "", false, 0, "")
	p.Ln(20)

	p.SetFont("Arial", "", 10)
	for _, e := range w.Suma.NeophodanMaterijal {
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
}

func (w *WingCal) specifikacijaRadovaList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()
	p.SetFont("Times", "B", 16)
	aktivnosti = fmt.Sprintf("Specifikacija radova %d/{nb}", p.PageNo())
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
}

func (w *WingCal) aktivnostiSuma(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, tr(w.text("Lista aktivnosti")), "0", 0, "", false, 0, "")
	p.Ln(20)
	p.SetFont("Arial", "", 10)
	for _, e := range w.Suma.Elementi {
		cols := []float64{40, pagew - mleft - mright - 20}
		rows := [][]string{
			[]string{
				e.Sifra, e.Element.Naziv,
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
	}

	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Suma: ")+fmt.Sprintf("%.2f", w.Suma.SumaCena), "0", 0, "", false, 0, "")
}

func (w *WingCal) specifikacijaMaterijalaList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Specifikacija materijala"), "0", 0, "", false, 0, "")
	materijal = fmt.Sprintf("Specifikacija materijala %d/{nb}", p.PageNo())
	p.Ln(20)

	p.SetFont("Arial", "", 10)
	for _, e := range w.Suma.NeophodanMaterijal {
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
	p.CellFormat(0, 10, w.text("Suma materijal: ")+fmt.Sprintf("%.2f", projekat.Elementi.SumaCena), "0", 0, "", false, 0, "")
	p.Ln(20)

}

func (w *WingCal) materijalSuma(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Lista materijala"), "0", 0, "", false, 0, "")
	p.Ln(20)
	p.SetFont("Arial", "", 10)
	for _, e := range w.Suma.NeophodanMaterijal {
		cols := []float64{40, pagew - mleft - mright - 20}
		//rows := [][]string{}
		rows := [][]string{
			[]string{
				fmt.Sprint(e.Id), e.Materijal.Naziv,
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
	}
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Suma materijal: ")+fmt.Sprintf("%.2f", projekat.Elementi.SumaCena), "0", 0, "", false, 0, "")
}

func (w *WingCal) sadrzajList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64) {
	p.AddPage()
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Sadržaj"), "0", 0, "", false, 0, "")
	p.Ln(20)
	p.CellFormat(0, 10, w.text(aktivnosti), "0", 0, "", false, 0, "")
	p.Ln(20)
	p.CellFormat(0, 10, w.text(materijal), "0", 0, "", false, 0, "")
	p.Ln(20)
}

func (w *WingCal) ponuda(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()
	p.SetFont("Times", "B", 18)
	p.CellFormat(0, 10, w.text("Ponuda"), "0", 0, "", false, 0, "")
	p.Ln(10)
	w.investitorList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
	p.Ln(20)
	w.aktivnostiSuma(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
	p.Ln(20)
	w.materijalSuma(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
}
func (w *WingCal) ipList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()
	w.projektantList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
	p.Ln(10)
	w.investitorList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
}
func (w *WingCal) investitorList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Investitor"), "0", 0, "", false, 0, "")
	p.Ln(10)
	p.SetFont("Arial", "", 10)
	cols := []float64{40, pagew - mleft - mright - 20}
	rows := [][]string{
		[]string{
			"MB", projekat.Investitor.MB,
		},
		[]string{
			"PIB", projekat.Investitor.PIB,
		},
		[]string{
			"KratakNaziv", projekat.Investitor.KratakNaziv,
		},
		[]string{
			"DugiNaziv", projekat.Investitor.DugiNaziv,
		},
		[]string{
			"Delatnost", projekat.Investitor.Delatnost,
		},
		[]string{
			"Adresa", projekat.Investitor.Adresa,
		},
		[]string{
			"Grad", projekat.Investitor.Grad,
		},
		[]string{
			"Email", projekat.Investitor.Email,
		},
		[]string{
			"DatumOsnivanja", projekat.Investitor.DatumOsnivanja,
		},
		//[]string{
		//	"Racuni", projekat.Investitor.Racuni,
		//},
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
}

func (w *WingCal) projektantList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("Nadzor"), "0", 0, "", false, 0, "")
	p.Ln(10)
	p.SetFont("Arial", "", 10)
	cols := []float64{40, pagew - mleft - mright - 20}
	rows := [][]string{
		[]string{
			"MB", projekat.Projektant.MB,
		},
		[]string{
			"PIB", projekat.Projektant.PIB,
		},
		[]string{
			"KratakNaziv", projekat.Projektant.KratakNaziv,
		},
		[]string{
			"DugiNaziv", projekat.Projektant.DugiNaziv,
		},
		[]string{
			"Delatnost", projekat.Projektant.Delatnost,
		},
		[]string{
			"Adresa", projekat.Projektant.Adresa,
		},
		[]string{
			"Grad", projekat.Projektant.Grad,
		},
		[]string{
			"Email", projekat.Projektant.Email,
		},
		[]string{
			"DatumOsnivanja", projekat.Projektant.DatumOsnivanja,
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
}
