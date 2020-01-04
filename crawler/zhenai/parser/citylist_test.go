package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T)  {
	//contents,err :=fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents,err := ioutil.ReadFile("citylist_data.html")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",contents)
}
