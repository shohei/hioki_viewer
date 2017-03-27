package main

import (
// "github.com/k0kubun/pp"
"./parser"
"fmt"
"net/http"
"encoding/json"
"math/rand"
"time"
"flag"
)

var ip *string

type DataArray struct {
	Number int
	SensorData *parser.SensorData
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://"+ *ip +"/REALDATA.HTM"

	t := time.Now()
	utc := t.UTC()
	year := utc.Year()
	month := int(utc.Month())
	day := utc.Day()
	hour := utc.Hour()
	minute := utc.Minute()
	second := utc.Second()
	// fmt.Println(year,month,day,hour,minute,second)
	datetime := []int {year,month,day,hour,minute,second}

	ds := parser.Parse(url,datetime)
	das := make([]DataArray,16)
	for i,d := range ds {
		da := &das[i]
		da.Number = i
		da.SensorData = d
	}
	bytes, _ := json.Marshal(das)

	// fmt.Println("Received Request")
	w.Header().Set("Content-Type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}

func main(){
	finish := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	ip = flag.String("ip", "192.168.100.210", "IP address of hioki logger")
	fmt.Println("server started.")

	staticServer := http.NewServeMux()
	staticServer.Handle("/",http.FileServer(http.Dir("./")))

	logServer := http.NewServeMux()
	logServer.HandleFunc("/",MainHandler)

	go func(){
		http.ListenAndServe(":8000",staticServer)
	}()
	go func(){
		http.ListenAndServe(":3000", logServer)
	}()
	<- finish
}
