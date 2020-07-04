package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"github.com/gioapp/gel"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	//post             = new(model.DuoCMSpost)
	stampajDugme  = new(widget.Clickable)
	materijalList = &layout.List{
		Axis: layout.Vertical,
	}
	putanjaList = &layout.List{
		Axis: layout.Vertical,
	}

	selected       int
	izbornikRadova = &layout.List{
		Axis: layout.Vertical,
	}
	elementOpis = &layout.List{
		Axis: layout.Vertical,
	}
	sumList = &layout.List{
		Axis: layout.Vertical,
	}
	neophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	ukupanNeophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	izborVrsteRadovaPanel = &layout.List{
		Axis: layout.Vertical,
	}
	dodajDugme          = new(widget.Clickable)
	nazadDugme          = new(widget.Clickable)
	zatvoriElementDugme = new(widget.Clickable)
	kolicina            = &gel.DuoUIcounter{
		Value:        1,
		OperateValue: 1,
		From:         1,
		To:           100,
		CounterInput: &widget.Editor{
			Alignment:  text.Middle,
			SingleLine: true,
		},
		CounterIncrease: new(widget.Clickable),
		CounterDecrease: new(widget.Clickable),
		CounterReset:    new(widget.Clickable),
	}
)

type WingCal struct {
	Naziv                    string
	Strana                   string
	LinkoviIzboraVrsteRadova map[int]*widget.Clickable
	EditPolja                *model.EditabilnaPoljaVrsteRadova
	Materijal                map[int]*model.WingMaterijal
	Radovi                   model.WingVrstaRadova
	Putanja                  []string
	IzbornikRadova           map[int]model.ElementMenu
	Transfered               model.WingCalGrupaRadova
	Client                   *model.Client
	PrikazaniElement         *model.WingVrstaRadova
	Suma                     *model.WingIzabraniElementi
	Podvrsta                 int
	Roditelj                 int
	Cyr                      bool
	Element                  bool
	UI                       WingUI
	API                      WingAPI
}

type WingUI struct {
	Device          string
	Window          *app.Window
	Tema            *gelook.WingUItheme
	Context         layout.Context
	Ekran           func(gtx layout.Context) layout.Dimensions
	D               layout.Dimensions
	C               layout.Context
	Ops             op.Ops
	BezMargine      layout.Inset
	SaMarginom      layout.Inset
	SaMalomMarginom layout.Inset
}

type WingAPI struct {
	OK     bool
	Adresa string
}
