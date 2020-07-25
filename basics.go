package main

import "fmt"

func main() {
	//arrays
	// array := [5]float64{1,2,3,4,5}
	//for loop with range
	// for _, value := range array {
	// 	fmt.Println(value)
	// } 
	// normal for loop
	// i := 1
	// for i<=10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//another type of normal for loop
	// for i :=0; i<=10; i++{
	// 	fmt.Println(i)
	// }

	//conditional statements
	// var myAge int = 18

	// switch {
	// case (myAge <18):
	// 	fmt.Println("You cannot drive")
	// 	fmt.Println("You cannot drive")
	// case (myAge>=18):
	// 	fmt.Println("You can drive")
	// }
	// if myAge < 18 {
	// 	fmt.Println("You cannot drive")
	// }else if myAge >= 18 {
	// 	fmt.Println("You can drive")
	// }

	//slices
	// numSlice := []int{1,2,3,4,5}
	// numSlice2 := []int{}
	// numSlice2 = numSlice[:]
	// numSlice2 = append(numSlice2,6,7,8,9,10,11)
	// for _, value := range numSlice {
	// 	fmt.Println(value)
	// }
	// for _, value := range numSlice2 {
	// 	fmt.Println(value)
	// }

	//call function with two return values
	// _, num2 := next2Values(5)
	// fmt.Println(num2)

	//`exceiption handling calling`
	// fmt.Println(div(2,0))
	// fmt.Println(div(2,2))
	// x := 1
	// changeValue(&x)
}
// function with two return values
// func next2Values(number int)(int,int)  {
// 	return number +1, number +2

// }

//exception handling	
// func div(num1,num2 int) int{
// 	defer func(){
// 		fmt.Println(recover())
// 	}()
// 	panic("Panic")
// 	return num1 / num2
// }
// pointers 
// func changeValue(x *int)  {
// 	*x = 10
// }