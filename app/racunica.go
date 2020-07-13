package calc

import (
	"github.com/w-ingsolutions/c/model"
)

func (w *WingCal) SumaRacunica() {
	w.SumaElementiPrikaz()
	s := 0.0
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena
	}
	w.Suma.SumaCena = s
	w.Suma.SumaCenaMaterijal, w.Suma.NeophodanMaterijal = w.NeopodanMaterijal(w.Suma.Elementi)
}

func (w *WingCal) PrikazaniElementSumaRacunica() func() {
	return func() {
		prikazaniElementSumaCena = w.PrikazaniElement.Cena * float64(w.UI.Counters.Kolicina.Value)
	}
}

func (w *WingCal) ProjekatRacunica() func() {
	return func() {
		p := w.Suma
		iz := []*model.WingIzabraniElement{}
		for _, e := range w.Suma.Elementi {
			ee := e
			ee.SumaCena = e.SumaCena * float64(w.UI.Counters.Radovi.Value) / 100
			iz = append(iz, e)
		}
		p.Elementi = iz
		projekat.Elementi = p
		//w.ProjekatSumaRacunica()
	}
}

func (w *WingCal) ProjekatMaterijalSumaRacunica() func() {
	return func() {
		projekat.Elementi.SumaCena, projekat.Elementi.NeophodanMaterijal = w.NeopodanMaterijal(projekat.Elementi.Elementi)
		//projekat.Elementi.SumaCena = w.Suma.SumaCenaMaterijal + w.Suma.SumaCenaMaterijal*float64(materijal.Value)/100
	}
}

func (w *WingCal) ProjekatSumaRacunica() func() {
	return func() {
		s := 0.0
		for _, e := range projekat.Elementi.Elementi {
			s = s + e.SumaCena
		}
		projekat.Elementi.SumaCena = s
	}
}
