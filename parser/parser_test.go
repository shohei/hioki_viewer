package parser 

import (
	"testing"
)

func TestParser(t *testing.T){
	url := "http://192.168.100.210/REALDATA.HTM"
	Parse(url)	
}

