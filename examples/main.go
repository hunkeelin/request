package main

import (
	"fmt"
	"github.com/hunkeelin/request"
	"io/ioutil"
)

func main() {
	rb := request.RequestBuilder{}
	resp, err := rb.SetURL("https://google.com").Do()
	if err != nil {
		panic(err)
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
}
