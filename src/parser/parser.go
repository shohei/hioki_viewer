package parser

import (
"github.com/k0kubun/pp"
"strconv"
"github.com/PuerkitoBio/goquery"
"strings"
"regexp"
)

type SensorData struct{
	name string
	datetime []int
	value float64
}

func ValueFloat(v string) float64{
	re := regexp.MustCompile("mV")
	r1 := re.ReplaceAllString(v,"")
	r2 := strings.Replace(r1, " ", "", -1)
	f,err := strconv.ParseFloat(r2,64)
	if (err!=nil){
		return 0.0	
	}
	return f
}

func Parse(url string, dt []int) []*SensorData{
	s := make([]*SensorData,20)
	pp.Print(s)

	dom,_ := goquery.NewDocument(url)
	dom.Find("table").Children().Each(func(i int, n *goquery.Selection){
		switch i {
		case 2:
			n.Children().Each(func(j int, nn *goquery.Selection){
				switch j {
				case 0:
					break
				case 1:
					break
				default:
					nn.Children().Each(func(k int, nnn *goquery.Selection){
						switch k {
						case 1:
							value_string := nnn.Text()
							s_num := 2*(j-1)-1
							s_name := "Sensor"+strconv.Itoa(s_num)
							ss := &SensorData{name: s_name, datetime: dt, value: ValueFloat(value_string)}

							s_idx := 2*(j-1)
							s[s_idx] = ss
						case 4:
							value_string := nnn.Text()
							s_num := 2*(j-1)
							s_name := "Sensor"+strconv.Itoa(s_num)
							ss := &SensorData{name: s_name, datetime: dt, value: ValueFloat(value_string)}

							s_idx := 2*(j-1)+1
							s[s_idx] = ss
						}
						})
				}
				})
		}
		})
	return s
}
