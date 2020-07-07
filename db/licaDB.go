package db

import (
	"github.com/w-ingsolutions/c/model"
)

func NewLica() (p []*model.WingLice, k []*model.WingLice) {
	for _, l := range lica() {
		if l.Pravno {
			k = append(k, l)
		} else {
			p = append(p, l)
		}
	}
	return
}
func lica() map[int]*model.WingLice {
	return map[int]*model.WingLice{
		0: &model.WingLice{
			Id:          0,
			Pravno:      false,
			Ime:         "Đorđe",
			Prezime:     "Marčetin",
			JMBG:        "11113101980034",
			BrojLicence: "U54mLj3n1H1dR4nT",
			Ulica:       "Josifa Marinkovića",
			Broj:        "24",
			Grad:        "Novi Sad",
			Admin:       true,
		},
		1: &model.WingLice{
			Id:          1,
			Pravno:      false,
			Ime:         "Čedomir",
			Prezime:     "Vukobrat",
			JMBG:        "111131019785555",
			BrojLicence: "XXXXXXXXXXXXX",
			Ulica:       "Iza Cao Picerije",
			Broj:        "44",
			Grad:        "Novi Sad",
			Admin:       true,
		},
		2: &model.WingLice{
			Id:     2,
			Pravno: true,
			Naziv:  "W-ing Solutions D.o.o.",
			Ulica:  "Bureval",
			Broj:   "244",
			Grad:   "Novi Sad",
			PIB:    "wingPIBPIBPIB",
			MB:     "wingMBMBMBMB",
		},
		3: &model.WingLice{
			Id:     3,
			Pravno: true,
			Naziv:  "Šuša",
			Ulica:  "Nedje Vjeternik",
			Broj:   "244",
			Grad:   "Fjutog",
			PIB:    "SSSSPIBPIBPIB",
			MB:     "SSSSMBMBMBMB",
		},
		4: &model.WingLice{
			Id:     4,
			Pravno: true,
			Naziv:  "Pošta koči čarda",
			Ulica:  "Ferenca Lajoša",
			Broj:   "332",
			Grad:   "Vranje",
			PIB:    "1212SSSSPIBPIBPIB",
			MB:     "1212SSSSMBMBMBMB",
		},
		5: &model.WingLice{
			Id:     5,
			Pravno: true,
			Naziv:  "Feritnica",
			Ulica:  "Magnetna polja",
			Broj:   "bb",
			Grad:   "Subotica",
			PIB:    "2323SSSSPIBPIBPIB",
			MB:     "2323SSSSMBMBMBMB",
		},
	}
}
