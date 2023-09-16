package main

import (
	"ecdownloader/internal"
	"ecdownloader/internal/handler"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	arguments := internal.Argument
	fmt.Println(arguments.Urls.Get())
	handler.ArgHandler(arguments)
}
