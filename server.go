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
	"os/exec"
	"regexp"
	"text/template"
)

var host_ip *string
var logger_ip *string

type Page struct {
	HostIp string 
}

type DataArray struct {
	Number int
	SensorData *parser.SensorData
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{HostIp: *host_ip}
	tpl, err1 := template.ParseFiles("index.html")
	if err1 != nil {
		panic(err1)
	}
	err2 := tpl.Execute(w,page)
	if err2 != nil {
		panic(err2)
	}
}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://"+ *logger_ip +"/REALDATA.HTM"

	t := time.Now()
	utc := t.UTC()
	year := utc.Year()
	month := int(utc.Month())
	day := utc.Day()
	hour := utc.Hour()
	minute := utc.Minute()
	second := utc.Second()
	datetime := []int {year,month,day,hour,minute,second}

	ds := parser.Parse(url,datetime)
	das := make([]DataArray,16)
	for i,d := range ds {
		da := &das[i]
		da.Number = i
		da.SensorData = d
	}
	bytes, _ := json.Marshal(das)

	w.Header().Set("Content-Type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
}

func main(){
	finish := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	cmdstr := "ifconfig | grep 192"
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if(err != nil){
		fmt.Println(err)
		panic(err)
	}
	outstr := string(out)
	r := regexp.MustCompile("192\\.\\d*\\.\\d*\\.\\d")
	res := r.FindAllString(outstr,-1)	
	host_ip := &res[0]
	fmt.Println("Host IP address:",*host_ip)

	logger_ip = flag.String("logger-ip", "192.168.100.210", "IP address of hioki logger")
	fmt.Println("Logger IP address:",*logger_ip)
	fmt.Println("server started.")

	staticServer := http.NewServeMux()
	staticServer.HandleFunc("/",staticHandler)

	logServer := http.NewServeMux()
	logServer.HandleFunc("/",MainHandler)

	go func(){
		http.ListenAndServe(":3000", logServer)
	}()
	go func(){
		http.ListenAndServe(":8000",staticServer)
	}()
	<- finish
}
