package parser

import (
"github.com/k0kubun/pp"
"github.com/PuerkitoBio/goquery"
"fmt"
)

func GetPage(url string){
	dom,_ := goquery.NewDocument(url)
	dom.Find("table").Children().Each(func(i int, n *goquery.Selection){
		switch i {
		case 2:
			n.Children().Each(func(j int, nn *goquery.Selection){
				switch j {
				case 0:
					pp.Print(nn.Text())
					fmt.Println()
				case 1:
					break
				default:
					nn.Children().Each(func(k int, nnn *goquery.Selection){
						switch k {
						case 0:
							fmt.Print("SENSOR: ")
							pp.Print(nnn.Text())
						case 1:
							pp.Print(nnn.Text())
							fmt.Println()
						case 3:
							fmt.Print("SENSOR: ")
							pp.Print(nnn.Text())
						case 4:
							pp.Print(nnn.Text())
							fmt.Println()
						}
						})
				}
				})
		}
		})
}


