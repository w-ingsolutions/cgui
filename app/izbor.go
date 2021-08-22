package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cgui/app/gel/container"
	"github.com/w-ingsolutions/cgui/app/gel/icontextbtn"
	"github.com/w-ingsolutions/cgui/app/gel/panel"
	"github.com/w-ingsolutions/cgui/app/helper"
	"github.com/w-ingsolutions/cgui/app/model"
)

var (
	IzborVrsteRadovaPanelElement = panel.NewPanel()
	Podvrstaradova               string
	Elementi                     string
)

func (w *WingCal) IzborPodVrsteRadova() func(gtx C) D {
	return func(gtx C) D {
		IzborVrsteRadovaPanelElement.PanelObject = w.IzbornikRadova
		IzborVrsteRadovaPanelElement.PanelObjectsNumber = len(w.IzbornikRadova)
		//izborVrsteRadovaPanel := panel.DuoUIpanelSt()
		//izborVrsteRadovaPanel.ScrollBar = w.UI.Tema.ScrollBarSt(0)
		//izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
		//return izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
		return izborVrsteRadovaPanel.Layout(gtx, len(w.IzbornikRadova), func(gtx C, i int) D {
			vrstarada := w.IzbornikRadova[i+1]
			//fmt.Println("vrstarada: ", vrstarada)
			//return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
			//	layout.Rigid(func(gtx C) D {
			//		return w.UI.BezMargine.Layout(gtx, func(gtx C) D {
			//layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
			//	layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
			//		layout.Rigid(func() {

			//btn := material.Button(w.UI.Tema.T, vrstarada.Link, fmt.Sprint(vrstarada.Id)+". "+w.text(vrstarada.Title))
			//icoNaziv := strings.Split(vrstarada.Title, " ")

			btn := icontextbtn.IconTextBtn(w.UI.Tema.T, vrstarada.Link, w.UI.Tema.Icons[vrstarada.Slug], unit.Dp(48), w.UI.Tema.Colors["Light"], fmt.Sprint(vrstarada.Id)+". "+w.text(vrstarada.Title))
			btn.CornerRadius = unit.Dp(0)
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["Gray"])
			if vrstarada.Materijal {
				btn.Background = helper.HexARGB(w.UI.Tema.Colors["DarkGray"])
			}

			w.LinkoviIzboraVrsteRadovaKlik(vrstarada)
			return btn.Layout(gtx)
		})
	}
}

func (w *WingCal) IzbornikRadovaStrana() func(gtx C) D {
	izbor := w.IzborVrsteRadova()
	if len(w.Putanja) > 1 {
		izbor = w.IzborPodVrsteRadova()
		if w.Element {
			izbor = w.PrikazaniElementIzgled()
		}
	}

	return func(gtx C) D {
		return layout.Inset{}.Layout(gtx, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X

			return lyt.Format(gtx, "vflexb(middle,r(_),r(_),f(1,_))",
				w.Izbornik(),
				//layout.Rigid(helper.Tema.DuoUIline(false, 4, 4, 0, w.UI.Tema.Colors["Dark"])),
				w.Nazad(),
				//layout.Rigid(helper.Tema.DuoUIline(false, 4, 4, 0, w.UI.Tema.Colors["Dark"])),
				izbor)
		})
	}
}

func (w *WingCal) Izbornik() func(gtx C) D {
	return func(gtx C) D {
		return putanjaList.Layout(gtx, len(w.Putanja), func(gtx C, i int) D {
			var p layout.Dimensions
			put := material.Body1(w.UI.Tema.T, w.text(w.Putanja[i].Title))
			put.Alignment = text.Middle
			if w.Putanja[i].Title != "Radovi" {
				return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
					return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						return put.Layout(gtx)
					})
				})
			}
			return p
		})
	}
}

//btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
//layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, w.Tema.Colors["Gray"])),

func (w *WingCal) LinkoviIzboraVrsteRadovaKlik(l model.ElementMenu) {
	for l.Link.Clicked() {
		fmt.Println("IZBOR", l.Title)

		switch len(w.Putanja) {
		case 1:
			//komanda = fmt.Sprint(l.Id)
			w.APIpozivIzbornik(w.Radovi.PodvrsteRadova[fmt.Sprint(l.Id+1)].PodvrsteRadova)
			Podvrstaradova = fmt.Sprint(l.Id)
			w.Podvrsta = l.Id
			fmt.Println("IZBOR 11111 ", l.Title)
			fmt.Println("IZBOR 11111 ", l.Id)
			fmt.Println("IZBOR w.Podvrsta ", w.Podvrsta)
			fmt.Println("IZBOR w.Roditelj ", w.Roditelj)
		case 2:
			//komanda = Podvrstaradova + "/" + fmt.Sprint(l.Id)
			fmt.Println("IZBOR 22222 ", l.Title)
			fmt.Println("IZBOR 22222 ", l.Id)
			fmt.Println("IZBOR w.Podvrsta ", w.Podvrsta)
			fmt.Println("IZBOR w.Roditelj ", w.Roditelj)
			w.APIpozivIzbornik(w.Radovi.PodvrsteRadova[fmt.Sprint(w.Podvrsta)].PodvrsteRadova[fmt.Sprint(l.Id)].PodvrsteRadova)

			Elementi = fmt.Sprint(l.Id)
			w.Roditelj = l.Id

		case 3:
			//komanda = Podvrstaradova + "/" + Elementi + "/" + fmt.Sprint(l.Id)
			fmt.Println("IZBOR 3333 ", l.Title)
			fmt.Println("IZBOR 3333 ", l.Id)
			fmt.Println("IZBOR w.Podvrsta ", w.Podvrsta)
			fmt.Println("IZBOR w.Roditelj ", w.Roditelj)
			w.APIpozivElement(w.Radovi.PodvrsteRadova[fmt.Sprint(w.Podvrsta)].PodvrsteRadova[fmt.Sprint(w.Roditelj)].PodvrsteRadova[fmt.Sprint(l.Id+1)])
			w.Element = true
		}
		if len(w.Putanja) < 3 {
			w.Putanja = append(w.Putanja, &model.ElementMenu{Id: l.Id + 1, Title: l.Title})
		}
		w.GenerisanjeLinkova(w.IzbornikRadova)
		w.UI.Counters.Kolicina.Value = 0
	}
}
