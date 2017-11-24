package main

import (
	"github.com/golang/tools/go/loader"
)

func main() {
	var conf loader.Config
	conf.Import("./testdata/")

	program, err := conf.Load()
	if err != nil {
		panic(err)
	}

	println("got program!")
	_ = program
}
