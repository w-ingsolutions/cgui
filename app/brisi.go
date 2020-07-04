package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/gelook"
)

func remove(slice []*model.WingIzabraniElement, s int) []*model.WingIzabraniElement {
	return append(slice[:s], slice[s+1:]...)
}

func (w WingCal) brisi(element *model.WingIzabraniElement, i int) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(w.UI.Tema.T, element.DugmeBrisanje, w.UI.Tema.Icons["Delete"])
		btn.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
		btn.Color = gelook.HexARGB(w.UI.Tema.Colors["Danger"])
		btn.Size = unit.Dp(16)
		btn.Background = gelook.HexARGB(w.UI.Tema.Colors["White"])
		for element.DugmeBrisanje.Clicked() {
			w.Suma.Elementi = remove(w.Suma.ElementiPrikaz, i)
			//w.Suma.Elementi = append(w.Suma.Elementi[:i], w.Suma.Elementi[i+1:]...)
			//w.Suma.Elementi[i] = nil
			//tabelaSuma = map[int]int{}
			w.NeopodanMaterijal()
			w.SumaRacunica()
		}
		return btn.Layout(gtx)
		//w.tabela(3, w.Context.Dimensions.Size.X)
	}
}
