package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	wapp "github.com/w-ingsolutions/c/pkg/app"
)

func NewWingCal() *WingCal {
	w := &WingCal{
		Dir:              wapp.Dir("wing", false),
		Naziv:            "W-ing Solutions - Kalkulator",
		Cyr:              true,
		Strana:           "radovi",
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:                 []*model.WingIzabraniElement{},
			UkupanNeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	w.NewMaterijal()
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
		Device:      "p",
		TopSpace:    28,
		BottomSpace: 56,
		Window: app.NewWindow(
			app.Size(unit.Dp(999), unit.Dp(1024)),
			app.Title("W-ing Solutions - Kalkulator"),
		),
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

	w.Putanja = append(w.Putanja, w.Radovi.Naziv)
	return w
}
