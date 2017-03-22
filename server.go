package main

import (
"./parser"
"fmt"
"net/http"
"github.com/zenazn/goji"
"github.com/zenazn/goji/web"
    // "reflect"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main(){
	fmt.Println("server started.")
	url := "http://192.168.100.210/REALDATA.HTM"
	parser.GetPage(url)
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
