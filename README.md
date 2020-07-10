# request 
[![CircleCI](https://circleci.com/gh/hunkeelin/request.svg?style=shield)](https://circleci.com/gh/hunkeelin/request)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunkeelin/request)](https://goreportcard.com/report/github.com/hunkeelin/request)
[![GoDoc](https://godoc.org/github.com/hunkeelin/request?status.svg)](https://godoc.org/github.com/hunkeelin/request)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hunkeelin/request/master/LICENSE)
## Motivations

This is a practice library to build request. The library is suppose to be minimalistic to avoid any type of bugs and unexpected behavior. 

## Golang version

`request` is currently compatible with golang version from 1.12+.

## Request Builder [![GoDoc](https://godoc.org/github.com/hunkeelin/request?status.svg)](https://godoc.org/github.com/hunkeelin/request#ReqBuilder)
```go
package main
  
import (
    "fmt"
    request "github.com/hunkeelin/request"
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
```

## CHANGELOG
- Added a way to make mtls request
