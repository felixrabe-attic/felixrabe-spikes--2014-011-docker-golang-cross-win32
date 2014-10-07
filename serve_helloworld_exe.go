package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var helloworld_exe string

func init() {
	c, err := ioutil.ReadFile("helloworld.exe")
	if err != nil {
		log.Fatal("Unable to read helloworld.exe")
	}
	helloworld_exe = string(c)
}

func serveHelloworld(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/octet-stream",
	)
	io.WriteString(
		res,
		helloworld_exe,
	)
}

func main() {
	http.HandleFunc("/", serveHelloworld)
	http.ListenAndServe(":9000", nil)
}
