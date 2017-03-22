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
	ch := make(chan int)

	url := "http://192.168.100.210/REALDATA.HTM"
	go parser.ParsePeriodically(url,ch)

	goji.Get("/hello/:name", hello)
	goji.Serve()
}
