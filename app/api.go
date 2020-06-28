package calc

import (
	"encoding/json"
	"fmt"
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (w *WingCal) GenerisanjeLinkova(radovi map[int]model.ElementMenu) {
	for rad, _ := range radovi {
		w.LinkoviIzboraVrsteRadova[rad] = new(widget.Clickable)
	}
	return
}

func (w *WingCal) GenerisanjeDugmicaBrisanje(radovi map[int]string) {
	for rad, _ := range radovi {
		w.LinkoviIzboraVrsteRadova[rad] = new(widget.Clickable)
	}
	return
}

func (w *WingCal) APIpozivIzbornik(komanda string) {
	radovi := map[int]model.ElementMenu{}
	api, err := w.APIpoziv(w.API.Adresa, komanda)
	if err != nil {
		w.API.OK = false
	} else {
		jsonErr := json.Unmarshal(api, &radovi)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.IzbornikRadova = radovi
	}
}
func (w *WingCal) APIpozivElementi(komanda string) {
	radovi := map[int]model.ElementMenu{}
	api, err := w.APIpoziv(w.API.Adresa, komanda)
	if err != nil {
		w.API.OK = false
	} else {
		jsonErr := json.Unmarshal(api, &radovi)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.IzbornikRadova = radovi
	}
}

func (w *WingCal) APIpozivElement(komanda string) {
	rad := &model.WingVrstaRadova{}
	api, err := w.APIpoziv(w.API.Adresa, komanda)
	if err != nil {
		w.API.OK = false
	} else {
		jsonErr := json.Unmarshal(api, &rad)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.PrikazaniElement = rad
	}
}

func (w *WingCal) APIpoziv(adresa, komanda string) ([]byte, error) {
	var body []byte
	url := adresa + komanda
	fmt.Println("url", url)
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "wing")
	res, err := spaceClient.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, err = ioutil.ReadAll(res.Body)
	}
	if err != nil {
		return nil, err
		//log.Fatal(readErr)
	}
	if body != nil {
		//defer body.Close()
	}
	return body, err
}
