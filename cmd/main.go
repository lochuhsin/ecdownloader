package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	intVal := 0
	flag.IntVar(&intVal, "name", 12345, "help message for flagname")
	stringVal := ""
	flag.StringVar(&stringVal, "name2", "fuck", "asdfasdf")
	flag.Parse()
	fmt.Println(intVal)
	fmt.Println(stringVal)
}
