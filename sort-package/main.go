package main

import (
	"fmt"
	"sort"
)


func main(){
	// vars :=[]int{5,2,0,3,4,9,6}
	// sort.Ints(vars)
	// fmt.Println(vars)

	str:=[]string{"Learning,"Golang","c"}
	sort.Strings(str)
	fmt.Println(str)

}