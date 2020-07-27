package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"

)

func main()  {
	file,err := os.Create("samp.txt")
	if err !=nil{
		log.Fatal(err)
	}
	file.WriteString("text in a file")
	file.Close()
	stream,err := ioutil.ReadFile("samp.txt") 
	if err !=nil{
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}