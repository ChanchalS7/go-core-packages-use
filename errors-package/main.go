package main

import (
	// "errors"
	"fmt"
)

/*
New() function
Errorf() function

*/

func process(i int)error{
	if i%2==0 {
		// return errors.New("Only odd number allowed")
		return fmt.Errorf("Only odd numbers allowed,go: %d",i)
	}
	return nil

}


func checkError(e error){
	if e!=nil{
		fmt.Println(e)
	}
	fmt.Println("Operation successful")
}

func main(){

	

	// err:= errors.New("Custom error occurred")
	// fmt.Println(err)


	err:=process(3)
	checkError(err)
	err=process(2)
	checkError(err)

}