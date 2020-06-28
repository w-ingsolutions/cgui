package main

import "net/http"

func main() {
	//http.Handle("/", http.FileServer(http.Dir("./html")))
	go panic(http.ListenAndServe(":9999", http.FileServer(http.Dir("cgui"))))
}
