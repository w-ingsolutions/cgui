package main

import (
	"embed"
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/w-ingsolutions/cgui/app/helper"
	"log"
	"os"

	"github.com/w-ingsolutions/cgui/app"
	"github.com/w-ingsolutions/cgui/cfg"
	in "github.com/w-ingsolutions/cgui/cfg/ini"
)

//go:embed jsondb/radovi/*
var jsonDBradovi embed.FS

//go:embed jsondb/podradovi/*
var jsonDBpodradovi embed.FS

func main() {
	w := calc.NewWingCal()
	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(w.Podesavanja.File)
	//w.APIimportFS(jsonDBpodradovi,jsonDBradovi)
	w.APIimportFS(jsonDBradovi, jsonDBpodradovi)

	//fmt.Println("w.Radovi: ", w.Radovi.PodvrsteRadova["26"])

	//fmt.Println(" PodvrsteRadovaPodvrsteRadova::::::::::: ", w.Radovi.PodvrsteRadova)

	w.APIpozivIzbornik(w.Radovi.PodvrsteRadova)
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
				helper.Fill(w.UI.Context, helper.HexARGB(w.UI.Tema.
					Colors["Light"]), w.UI.Context.Constraints.Max)

				//if !w.API.OK {
				//	w.GreskaEkran()
				//} else {
				w.GlavniEkran()
				//}

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}
