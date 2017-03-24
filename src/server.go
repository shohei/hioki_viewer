package main

import (
"./parser"
"fmt"
"net/http"
"encoding/json"
"math/rand"
"time"
"flag"
	// "reflect"
)

type Data struct {
	Name         string    `json:"name"`
	CreatedAt    []int `json:"created_at"`
	Value        int       `json:"value"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ip := flag.String("ip", "192.168.100.210", "IP address of hioki logger")
	url := "http://"+ *ip +"/REALDATA.HTM"

	t := time.Now()
	utc := t.UTC()
	year := utc.Year()
	month := int(utc.Month())
	day := utc.Day()
	hour := utc.Hour()
	minute := utc.Minute()
	second := utc.Second()
	fmt.Println(year,month,day,hour,minute,second)
	datetime := []int {year,month,day,hour,minute,second}

	d := parser.Parse(url,datetime)
	bytes, _ := json.Marshal(d)

	fmt.Println("Received Request")
	w.Header().Set("Content-Type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
	fmt.Println(bytes)
}

func main(){
	rand.Seed(time.Now().UnixNano())

	fmt.Println("server started.")
	http.HandleFunc("/",MainHandler)
	http.ListenAndServe(":3000", nil)
}
