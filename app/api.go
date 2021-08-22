package calc

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/w-ingsolutions/cgui/app/model"
	"log"
	"path/filepath"
	"strconv"
)

func (w *WingCal) APIimportFS(jsonDBradovi, jsonDBpodradovi embed.FS) {
	w.Radovi = &model.WingVrstaRadova{
		Id:             0,
		Naziv:          "Radovi",
		Opis:           "Radovi predmer",
		PodvrsteRadova: make(map[string]*model.WingVrstaRadova),
	}
	radoviFolder, err := jsonDBradovi.ReadDir(filepath.Join(w.Podesavanja.Path, "radovi"))
	if err != nil {
	}
	for _, podRadoviRaw := range radoviFolder {
		podvrstaRadova := &model.WingVrstaRadova{}
		pod, err := jsonDBradovi.ReadFile(filepath.Join(w.Podesavanja.Path, "radovi", podRadoviRaw.Name()))
		if err != nil {
			fmt.Println("Err", err)
		}
		jsonErr := json.Unmarshal(pod, &podvrstaRadova)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.Radovi.PodvrsteRadova[strconv.Itoa(podvrstaRadova.Id)] = podvrstaRadova
	}

	podradoviFolder, err := jsonDBpodradovi.ReadDir(filepath.Join(w.Podesavanja.Path, "podradovi"))

	for _, podradoviFolderRaw := range podradoviFolder {
		//vrstaRadova := &model.WingVrstaRadova{
		//	Id:             0,
		//	Naziv:          apiFolderRaw.Name(),
		//	PodvrsteRadova: make(map[string]*model.WingVrstaRadova),
		//}
		w.Radovi.PodvrsteRadova[podradoviFolderRaw.Name()].PodvrsteRadova = make(map[string]*model.WingVrstaRadova)
		apiFolder, err := jsonDBpodradovi.ReadDir(filepath.Join(w.Podesavanja.Path, "podradovi", podradoviFolderRaw.Name()))
		if err != nil {
		}
		fmt.Println("podradoviFolderRaw.Name: ", podradoviFolderRaw.Name())

		for _, podRadoviRaw := range apiFolder {
			podvrstaRadova := &model.WingVrstaRadova{}
			pod, err := jsonDBpodradovi.ReadFile(filepath.Join(w.Podesavanja.Path, "podradovi", podradoviFolderRaw.Name(), podRadoviRaw.Name()))
			if err != nil {
				fmt.Println("Err", err)
			}
			jsonErr := json.Unmarshal(pod, &podvrstaRadova)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			//fmt.Println("pod: ", pod)
			w.Radovi.PodvrsteRadova[podradoviFolderRaw.Name()].PodvrsteRadova[strconv.Itoa(podvrstaRadova.Id)] = podvrstaRadova
		}
	}
	return
}

func (w *WingCal) APIpozivIzbornik(podvrsteRadova map[string]*model.WingVrstaRadova) {
	radovi := map[int]model.ElementMenu{}
	for _, podRadovi := range podvrsteRadova {
		menuRadovi := model.ElementMenu{
			Id:    podRadovi.Id,
			Title: podRadovi.Naziv,
			Slug:  podRadovi.Slug,
		}
		radovi[menuRadovi.Id] = menuRadovi
	}
	w.IzbornikRadova = radovi
	return
}

func (w *WingCal) APIpozivElementi(komanda string) {
	radovi := map[int]model.ElementMenu{}
	for _, podRad := range w.Radovi.PodvrsteRadova[komanda].PodvrsteRadova {
		fmt.Println("APIpozivElementi", podRad)
		w.IzbornikRadova = radovi
	}
}

func (w *WingCal) APIpozivElement(podvrstaRadova *model.WingVrstaRadova) {
	fmt.Println("IZBOR podvrstaRadova 0", podvrstaRadova.Naziv)
	fmt.Println("IZBOR Putanja 1", w.Putanja[1].Title)
	//fmt.Println("IZBOR Element ", w.Radovi.PodvrsteRadova[fmt.Sprint(w.Putanja[1][0]-1)])
	//fmt.Println("IZBOR Element ", w.Radovi.PodvrsteRadova[w.Putanja[0][0]].Naziv)
	//fmt.Println("IZBOR Element ", w.Radovi.PodvrsteRadova[w.Putanja[0][0]].PodvrsteRadova[w.Putanja[1][0]])
	//fmt.Println("IZBOR Element ", w.Radovi.PodvrsteRadova[w.Putanja[0][0]].PodvrsteRadova[w.Putanja[1][0]].Naziv)
	//fmt.Println("IZBOR Element ", l.Id)
	//fmt.Println("IZBOR w.Podvrsta ", w.Podvrsta)

	//rad := &model.WingVrstaRadova{}
	//api, err := jsonDBpodradovi.ReadFile(filepath.Join(w.Podesavanja.Path, komanda))
	//if err != nil {
	//w.API.OK = false
	//} else {
	//	jsonErr := json.Unmarshal(api, &rad)
	//	if jsonErr != nil {
	//		log.Fatal(jsonErr)
	//	}
	//	fmt.Println("APIpozivElementtttt", api)
	//
	w.PrikazaniElement = podvrstaRadova
	//}
}

//
//
//func (w *WingCal) APIpozivIzbornik(komanda string) {
//	radovi := map[int]model.ElementMenu{}
//	api, err := w.APIpoziv(w.API.Adresa, komanda)
//	if err != nil {
//		w.API.OK = false
//	} else {
//		jsonErr := json.Unmarshal(api, &radovi)
//		if jsonErr != nil {
//			log.Fatal(jsonErr)
//		}
//		w.IzbornikRadova = radovi
//	}
//	fmt.Println("radoviradoviradovi", radovi)
//
//}
//func (w *WingCal) APIpozivElementi(komanda string) {
//	radovi := map[int]model.ElementMenu{}
//	api, err := w.APIpoziv(w.API.Adresa, komanda)
//	if err != nil {
//		w.API.OK = false
//	} else {
//		jsonErr := json.Unmarshal(api, &radovi)
//		if jsonErr != nil {
//			log.Fatal(jsonErr)
//		}
//		w.IzbornikRadova = radovi
//	}
//}
//
//func (w *WingCal) APIpozivElement(komanda string) {
//	rad := &model.WingVrstaRadova{}
//	api, err := w.APIpoziv(w.API.Adresa, komanda)
//	if err != nil {
//		w.API.OK = false
//	} else {
//		jsonErr := json.Unmarshal(api, &rad)
//		if jsonErr != nil {
//			log.Fatal(jsonErr)
//		}
//		w.PrikazaniElement = rad
//	}
//}

//func (w *WingCal) APIpoziv(adresa, komanda string) ([]byte, error) {
//	var body []byte
//	url := adresa + komanda
//	fmt.Println("url", url)
//	spaceClient := http.Client{
//		Timeout: time.Second * 10, // Maximum of 2 secs
//	}
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("User-Agent", "wing")
//	res, err := spaceClient.Do(req)
//	if err != nil {
//		return nil, err
//	} else {
//		body, err = ioutil.ReadAll(res.Body)
//	}
//	if err != nil {
//		return nil, err
//		//log.Fatal(readErr)
//	}
//	if body != nil {
//		//defer body.Close()
//	}
//	return body, err
//}
