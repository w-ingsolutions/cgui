package in

import (
	"encoding/json"
	"fmt"
	"github.com/w-ingsolutions/cgui/cfg"
	"io/ioutil"
	"os"

	"git.parallelcoin.io/dev/utl"
)

// Init checks to see if an appdata directory has been created
// and if it hasn't, it expects the bitnodes credentials in the
// CLI parameters which it writes to the config
func Init() {
	fmt.Println("initializing")

	if !utl.FileExists(cfg.File) {
		// cfguration does not exist, so we expect arg 1 and 2 to be
		// the username and password for bitnodes

		if !utl.EnsureDir(cfg.File) {
			panic("could not create data directory to write config")
		}
		//os.Mkdir(cfg.Dir+"/www", 0700)
		// if !util.EnsureDir(filepath.Join(Coins, "something")) {
		// 	panic("could not create coins directory")
		// }
		fh, err := os.Create(cfg.File)
		if err != nil {
			panic("could not create config file:" + err.Error())
		}
		j, e := json.MarshalIndent(cfg.Credentials, "", "\t")
		if e != nil {
			panic("JSON marshalling error:" + e.Error())
		}
		_, err = fmt.Fprint(fh, string(j))
		if err != nil {
			panic("error writing configuration to file:" + err.Error())
		}
	} else {
		// If the cfg file exists, read it in
		confText, err := ioutil.ReadFile(cfg.File)
		if err != nil {
			panic(err)
		}
		e := json.Unmarshal(confText, &cfg.Credentials)
		if e != nil {
			panic(e)
		}
	}
}
