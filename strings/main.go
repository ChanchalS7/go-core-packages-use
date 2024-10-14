package main

import (
	"fmt"
	"strings"
)

//Contains
/*
func Contains(str string, substr string) bool
*/

//ReplaceAll
/*
func ReplaceAll(str, old, new string) string
*/

//Count
/*
func Count(first,second string)int
first- we have to search
second- we need to search
*/


func main(){
	// text := "Go is fun"
	// fmt.Println(text)
	// text[1]='l' //strings are immutable

	// learning := "learning standard library in Go"
	// fun := "library in Go" //learning in Go - false (always look for exact match)

	// result := strings.Contains(learning,fun)
	// fmt.Println(result)

	// learning:= "learning standard library in python learning"
	// fun:="learning"

	// result:=strings.Count(learning,fun)
	// fmt.Println(result)

learning:= "learning standard library in Python"
	fun:="library in Python"
	result:=strings.ReplaceAll(learning,fun,"library in Go")
	fmt.Println(result)
}