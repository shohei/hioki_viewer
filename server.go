package main

import (
	// "./parser"
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	// "github.com/zenazn/goji"
	// "github.com/zenazn/goji/web"
	// "reflect"
)

// func hello(c web.C, w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
// }

var counter func() int = increment()

// type Marshaler interface {
//     MarshalJSON() ([]byte, error)
// }
//
// type JSONTime time.Time
//
// func (t JSONTime)MarshalJSON() ([]byte, error) {
//     stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
//     return []byte(stamp), nil
// }
//
// type Document struct {
//     Name				string
//     Value       int 
//     Stamp       JSONTime
// }

type Data struct {
  Name         string    `json:"name"`
  Value        int       `json:"value"`
  CreatedAt time.Time `json:"created_at"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// i := counter()
	// i := counter()
	// numbers := []int{i,a}
	// js, err := json.Marshal(numbers)
  d := Data{
    "Sensor1",
    rand.Intn(10),
		time.Now(),
  }
  bytes, _ := json.Marshal(d)

	// doc := Document{"Sensor1", a, JSONTime(time.Now())}   
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	fmt.Println("Received Request")
	w.Header().Set("Content-Type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
	fmt.Println(bytes)
}

func increment() func() int{
	i := 0
	return func() int{
	  i++
		return i
	}
}

func main(){
	// ch := make(chan int)
	// url := "http://192.168.100.210/REALDATA.HTM"
	// go parser.ParsePeriodically(url,ch)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("server started.")
	http.HandleFunc("/",MainHandler)
	http.ListenAndServe(":3000", nil)

	// goji.Get("/hello/:name", hello)
	// goji.Serve()
}
