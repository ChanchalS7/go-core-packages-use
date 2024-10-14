package main

import (
	"fmt"
	"strings"
)
func main(){
	//io.Reader method internally
	//use inbuilt NewReader method
	r :=strings.NewReader("Learning is fun")
	//create slice fo size 4
	buf:=make([]byte,4)

	//1.
	// n,err:=r.Read(buf)
	// fmt.Println(n,err)
	//2.
	for{
		n,err:=r.Read(buf)
	// fmt.Println(n,err)
	// fmt.Println(buf[:n],err)
	fmt.Println(string(buf[:n]),err)
	if err!=nil{
		fmt.Println("breaking out")
		break
	}
	}
}