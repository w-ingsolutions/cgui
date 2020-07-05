package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
	"github.com/w-ingsolutions/c/pkg/homedir"
)

func NewWingCal() *WingCal {
	home, err := homedir.UserHomeDir()
	if err != nil {
	}
	w := &WingCal{
		Dir:   home + "/wing",
		Naziv: "W-ing Solutions - Kalkulator",

		Strana:           "izbornik",
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:                 []*model.WingIzabraniElement{},
			UkupanNeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	w.NewMaterijal()
	w.Radovi = model.WingVrstaRadova{
		Id:       0,
		Naziv:    "Radovi",
		Slug:     "radovi",
		Omogucen: false,
		Baza:     false,
		Element:  false,
	}

	w.UI = WingUI{
		Device: "p",
		Window: app.NewWindow(
			app.Size(unit.Dp(999), unit.Dp(1024)),
			app.Title("W-ing Solutions - Kalkulator"),
		),
		Tema:            gelook.NewWingUItheme(),
		BezMargine:      layout.UniformInset(unit.Dp(0)),
		SaMarginom:      layout.UniformInset(unit.Dp(8)),
		SaMalomMarginom: layout.UniformInset(unit.Dp(4)),
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "https://wing.marcetin.com/",
	}

	w.Putanja = append(w.Putanja, w.Radovi.Naziv)
	return w
}
