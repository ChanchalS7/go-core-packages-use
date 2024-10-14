package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	r:= strings.NewReader("Some io.Reader stream to be read\n")

	if _, err:= io.Copy(os.Stdout,r); err!=nil{
		log.Fatal(err)
	}
}

//go doc io.Copy
//go doc os.Stdout