package main

import (
	"go/parser"

	"github.com/golang/tools/go/loader"
)

func main() {
	var conf loader.Config
	conf.Import("./testdata/")
	conf.ParserMode = parser.ParseComments

	program, err := conf.Load()
	if err != nil {
		panic(err)
	}

	println("got program!")
	_ = program
}
