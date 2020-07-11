package calc

func (w *WingCal) SumaRacunica() {
	w.SumaElementiPrikaz()
	s := 0.0
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena
	}
	w.Suma.SumaCena = s
}

func (w *WingCal) ProjekatSumaRacunica() func() {
	return func() {
		projekat.Radovi = w.Suma.Elementi
	}
}

func (w *WingCal) PrikazaniElementSumaRacunica() func() {
	return func() {
		prikazaniElementSumaCena = w.PrikazaniElement.Cena * float64(kolicina.Value)
	}
}
