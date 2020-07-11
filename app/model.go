package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/cache"
	"github.com/w-ingsolutions/c/pkg/translate"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	//post             = new(model.DuoCMSpost)
	prikazaniElementSumaCena float64
	selected                 int
	projekat                 = new(model.WingProjekat)
	latCyrBool               = new(widget.Bool)

	projektantIzbor = new(widget.Enum)
	klijentiIzbor   = new(widget.Enum)

	kolicina = &counter.DuoUIcounter{
		Value:        1,
		OperateValue: 1,
		From:         1,
		To:           999,
		CounterInput: &widget.Editor{
			Alignment:  text.Middle,
			SingleLine: true,
			Submit:     true,
		},

		CounterIncrease: new(widget.Clickable),
		CounterDecrease: new(widget.Clickable),
		CounterReset:    new(widget.Clickable),
	}
	radovi = &counter.DuoUIcounter{
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
	materijal = &counter.DuoUIcounter{
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
	Strana string
	//LinkoviIzboraVrsteRadova map[int]*widget.Clickable
	EditPolja        *model.EditabilnaPoljaVrsteRadova
	Materijal        map[int]*model.WingMaterijal
	Lica             WingUloge
	Radovi           model.WingVrstaRadova
	Putanja          []string
	IzbornikRadova   map[int]model.ElementMenu
	Transfered       model.WingCalGrupaRadova
	Client           *model.Client
	PrikazaniElement *model.WingVrstaRadova
	Suma             *model.WingIzabraniElementi
	Podvrsta         int
	Roditelj         int
	Element          bool
	UI               WingUI
	API              WingAPI
	Jezik            WingJezik
	Kes              cache.Cache
	Podesavanja      *WingPodesavanja
}

type WingUI struct {
	Device           string
	TopSpace         int
	BottomSpace      int
	Window           *app.Window
	Tema             *theme.DuoUItheme
	Context          layout.Context
	Ekran            func(gtx layout.Context) layout.Dimensions
	D                layout.Dimensions
	C                layout.Context
	Ops              op.Ops
	BezMargine       layout.Inset
	SaMarginom       layout.Inset
	SaMalomMarginom  layout.Inset
	SaStraneMarginom layout.Inset
}

type WingAPI struct {
	OK     bool
	Adresa string
}

type WingJezik struct {
	t        translate.Translation
	izabrani string
	dostupni []string
	linkovi  map[string]*widget.Clickable
}

type WingPodesavanja struct {
	Naziv string
	Dir   string
	File  string
	Cyr   bool
}

type WingUloge struct {
	Projektanti []*model.WingLice
	Investotori []*model.WingLice
}
