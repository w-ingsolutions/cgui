package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/pkg/gelook"
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
				gelook.WingUIfill(w.UI.Context, w.UI.Tema.Colors["Light"])

				if !w.API.OK {
					w.GreskaEkran()
				} else {
					w.GlavniEkran(w.UI.Context)
				}

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}
