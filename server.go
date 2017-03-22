package main

import (
    "fmt"
    "golang.org/x/net/html"
    "github.com/PuerkitoBio/goquery"
)

func showData(index int, s *goquery.Selection){
	count := 0
	func (){
		fmt.Println(count)
		fmt.Println(s.Children())
		// fmt.Println(s.Children()[0].Text())
		// fmt.Println(s.Children()[1].Text())
		// fmt.Println(s.Children()[3].Text())
		// fmt.Println(s.Children()[4].Text())
	}()
    count++
}

func GetPage(url string){
	dom,_ := goquery.NewDocument(url)
	fmt.Println(dom.Find("table").Nodes).Each(func(_ int, n *html.Node){
			fmt.Println(*n.Data)
	})
	// fmt.Println(dom.Find("table").Children().Text())
	// fmt.Println(dom.Find("table").Next().Text())
	// fmt.Println(dom.Find("table").Next().Next().Text())
// 	dom.Find("table").Next().Next().firstChild.Each(func(index int, s *goquery.Selection){
// 		switch(index){
// 			case 0:
// 				fmt.Println(s)
// 				// s.FirstChild
// 				fmt.Println("")
// 			case 1:
// 				return
// 			default:
// 				showData(index,s)
// 		}
// 	})
}

func main(){
  fmt.Println("server started.")
  url := "http://192.168.100.210/REALDATA.HTM"
  GetPage(url)
}
