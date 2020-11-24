package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// passing file directly
	data, err := ioutil.ReadFile("./assn_text_file.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	// getting file from command-line args
	arg := os.Args[1]
	data2, err2 := ioutil.ReadFile(arg)
	if err2 != nil {
		fmt.Println("Error 2: ", err2)
	}
	fmt.Println("from passed args: ", string(data2))
}
