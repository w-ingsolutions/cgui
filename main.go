package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/cgui/app"
	"log"
	"os"
)

func main() {
	w := calc.NewWingCal()
	w.LinkoviIzboraVrsteRadova = map[int]*widget.Clickable{}
	w.APIpozivIzbornik("radovi")

	w.GenerisanjeLinkova(w.IzbornikRadova)

	go func() {
		defer os.Exit(0)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *calc.WingCal) error {
	for {
		select {
		case e := <-w.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				w.UI.Context = layout.NewContext(&w.UI.Ops, e)

				//gelook.DuoUIfill(&w.Context, w.Tema.Colors["Light"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(w.UI.Context,
					layout.Rigid(
						material.H3(w.UI.Tema.T, "W-ing Solutions ").Layout),
					layout.Flexed(1, w.GlavniEkran()),
				)
				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}
