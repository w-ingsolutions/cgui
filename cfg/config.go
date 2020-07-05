package cfg

import (
	"github.com/w-ingsolutions/c/pkg/app"
	"path/filepath"
)

// Conf is the configuration for accessing bitnodes endpoint
type Conf struct {
	Username, Password, CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey string
}

// configurations for jorm
var (
	home = app.Dir("wing", false)

	File = filepath.Join(home, "conf.json")
	// Web is a subfolder because otherwise the config above would be served by the http.Dir webserver
	TSL         = home + "/tsl/"
	Web         = "/www/"
	Credentials = Conf{}
)
