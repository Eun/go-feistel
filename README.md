# go-feistel [![Travis](https://img.shields.io/travis/Eun/go-feistel.svg)](https://travis-ci.org/Eun/go-feistel) [![Codecov](https://img.shields.io/codecov/c/github/Eun/go-feistel.svg)](https://codecov.io/gh/Eun/go-feistel) [![GoDoc](https://godoc.org/github.com/Eun/go-feistel?status.svg)](https://godoc.org/github.com/Eun/go-feistel) [![go-report](https://goreportcard.com/badge/github.com/Eun/go-feistel)](https://goreportcard.com/report/github.com/Eun/go-feistel)
This package implements the Feistel cipher in go.  
It contains the pure Feistel [Encrypt](feistel.go#L5)/[Decrypt](feistel.go#L23) functions as well as
ECB and CBC implementations. 

```go
package main

import (
	"log"
	"github.com/Eun/go-feistel"
)

func main() {
	rounds := 2
	keys :=  []uint32{0xDEADBEEF, 0xFEEDFACE}
	secretText, err := feistel.ECB.Encrypt([]byte("Hello World"), rounds, keys)
	if err != nil {
		log.Fatal(err)
	}
	
	plainText, err := feistel.ECB.Decrypt(secretText, rounds, keys)
	if err != nil {
   		log.Fatal(err)
   	}
	
	log.Println(string(plainText))
}

```