package model

import (
	"gioui.org/widget"
)

type WingVrstaRadova struct {
	Id                 int                            `json:"id"`
	Naziv              string                         `json:"naziv"`
	Opis               string                         `json:"opis"`
	Obracun            string                         `json:"obracun"`
	Jedinica           string                         `json:"jedinica"`
	Cena               float64                        `json:"cena"`
	Slug               string                         `json:"slug"`
	Omogucen           bool                           `json:"omogucen"`
	Baza               bool                           `json:"baza"`
	Element            bool                           `json:"element"`
	PodvrsteRadova     map[string]*WingVrstaRadova    `json:"podvrsteradova"`
	NeophodanMaterijal map[int]WingNeophodanMaterijal `json:"neophodanmaterijal"`
}

type WingIzabraniElementi struct {
	Id                       string
	SumaCena                 float64
	SumaCenaMaterijal        float64
	Elementi                 []*WingIzabraniElement
	NeophodanMaterijal       map[int]WingNeophodanMaterijal
	NeophodanMaterijalPrikaz map[int]WingNeophodanMaterijal
}

type WingIzabraniElement struct {
	Sifra         string
	Kolicina      int
	SumaCena      float64
	Element       WingVrstaRadova
	DugmeBrisanje *widget.Clickable
}

type WingCalGrupaRadova struct {
	Id       string
	Slug     string
	Elementi map[int]WingVrstaRadova `json:"elementi"`
}
