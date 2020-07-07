package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	wapp "github.com/w-ingsolutions/c/pkg/app"
	"github.com/w-ingsolutions/c/pkg/latcyr"
	"github.com/w-ingsolutions/c/pkg/translate"
	"github.com/w-ingsolutions/cgui/db"
	"path/filepath"
)

func NewWingCal() *WingCal {
	w := &WingCal{
		Strana:           "radovi",
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:                 []*model.WingIzabraniElement{},
			UkupanNeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	w.Podesavanja = &WingPodesavanja{
		Naziv: "Kalkulator",
		Dir:   wapp.Dir("wing", false),
		Cyr:   true,
	}
	w.Podesavanja.File = filepath.Join(w.Podesavanja.Dir, "conf.json")
	w.Materijal = db.NewMaterijal()
	projektanti, klijenti := db.NewLica()
	w.Lica.Projektanti = projektanti
	w.Lica.Klijenti = klijenti
	saStraneMarginom := layout.UniformInset(unit.Dp(0))
	saStraneMarginom.Left = unit.Dp(8)
	saStraneMarginom.Right = unit.Dp(8)
	w.Radovi = model.WingVrstaRadova{
		Id:       0,
		Naziv:    "Radovi",
		Slug:     "radovi",
		Omogucen: false,
		Baza:     false,
		Element:  false,
	}
	w.UI = WingUI{
		Device:           "p",
		TopSpace:         28,
		BottomSpace:      56,
		Tema:             theme.NewDuoUItheme(),
		BezMargine:       layout.UniformInset(unit.Dp(0)),
		SaMarginom:       layout.UniformInset(unit.Dp(8)),
		SaMalomMarginom:  layout.UniformInset(unit.Dp(4)),
		SaStraneMarginom: saStraneMarginom,
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "https://wing.marcetin.com/",
	}
	w.Jezik = WingJezik{
		t:        translate.Translation{"sr", "sr"},
		dostupni: []string{"en", "ru", "de", "sl", "gr", "zh", "jp", "it"},
		linkovi:  make(map[string]*widget.Clickable),
	}
	w.GenerisanjeDugmicaJezici()
	w.Kes = make(map[string]string)

	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(999), unit.Dp(1024)),
		app.Title("W-ing "+w.Jezik.t.T(latcyr.C(w.Podesavanja.Naziv, w.Podesavanja.Cyr))),
	)

	w.Putanja = append(w.Putanja, w.Radovi.Naziv)
	return w
}
