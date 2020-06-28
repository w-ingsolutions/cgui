package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
)

func NewWingCal() *WingCal {
	wing := &WingCal{
		Naziv: "W-ing Solutions - Kalkulator",

		Strana:           "kalkulator",
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:                 []*model.WingIzabraniElement{},
			UkupanNeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	wing.NewMaterijal()
	wing.Radovi = model.WingVrstaRadova{
		Id:       0,
		Naziv:    "Radovi",
		Slug:     "radovi",
		Omogucen: false,
		Baza:     false,
		Element:  false,
	}

	wing.UI = WingUI{
		Window: app.NewWindow(
			app.Size(unit.Dp(1280), unit.Dp(1024)),
			app.Title("W-ing Solutions - Kalkulator"),
		),
		Tema:       gelook.NewWingUItheme(),
		BezMargine: layout.UniformInset(unit.Dp(0)),
		SaMarginom: layout.UniformInset(unit.Dp(8)),
	}
	wing.Putanja = append(wing.Putanja, wing.Radovi.Naziv)
	return wing
}
