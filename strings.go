package main

import (
	"fmt"
	"strings"
	"sort"
)

func main()  {
	// _string := "helloworld !"
	// fmt.Println(strings.Contains(_string,"l"))
	// fmt.Println(strings.Index(_string,"!"))
	// fmt.Println(strings.Count(_string,"l"))
	// fmt.Println(strings.Replace(_string,"hello","hi",1))

	csvString := "1,2,3,4,5"
	fmt.Println(strings.Split(csvString,","))

	listOfLetter := []string{"c","a","b"}
	sort.Strings(listOfLetter)
	fmt.Println(listOfLetter)

	listOfNums := strings.Join([]string{"c","a","b"},", ")

	fmt.Println(listOfNums)
}