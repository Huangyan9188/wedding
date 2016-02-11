package main

import (
	//"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	//http.HandleFunc("/testchangePwd", sayhello)
	http.ListenAndServe("0.0.0.0:8004", nil)
}
