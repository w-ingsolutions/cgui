package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
)

func NewWingCal() *WingCal {
	w := &WingCal{
		Naziv: "W-ing Solutions - Kalkulator",

		Strana:           "kalkulator",
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
		Window: app.NewWindow(
			app.Size(unit.Dp(1280), unit.Dp(1024)),
			app.Title("W-ing Solutions - Kalkulator"),
		),
		Tema:       gelook.NewWingUItheme(),
		BezMargine: layout.UniformInset(unit.Dp(0)),
		SaMarginom: layout.UniformInset(unit.Dp(8)),
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "http://212.62.35.158:9909/",
	}

	w.Putanja = append(w.Putanja, w.Radovi.Naziv)
	return w
}
