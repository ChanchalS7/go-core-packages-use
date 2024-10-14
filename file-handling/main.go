package main

import (
	"fmt"
	"os"
	// "path/filepath"
)


func main(){

	//path/filepath package
	// path:=filepath.Join("dir1","dir2","text.text") 
	// dir1,dir2/../dir3.text.text
	// fmt.Println(path)
	// fmt.Println(filepath.Dir(path))
	// fmt.Println(filepath.Base(path))
	// fmt.Println(filepath.IsAbs(path))
	// fmt.Println(filepath.IsAbs("/dir/file"))
	// fmt.Println(filepath.Ext(path))




	//OS package
	// fileInfo, err:= os.Stat("/Users/chanchalverma/learning/core-packages-go/file-handling/temp.txt")
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(fileInfo.Name(),fileInfo.Size(),fileInfo.Mode(),fileInfo.IsDir())


	// path:="/Users/chanchalverma/learning/core-packages-go/file-handling/temp.txt"
	// data,err:= os.ReadFile(path)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)
	// fmt.Println(string(data))

	// we read the file in chunk when file size is large

	// file, err := os.Open(path)
	// b:=make([]byte,4)
	// if err!=nil{
	// 	fmt.Println("Error:",err)
	// }
	// for{
	// 	n,err:=file.Read(b)
	// 	if err!=nil{
	// 		fmt.Println("Error:",err)
	// 		break
	// 	}
	// 	// fmt.Println(b[0:n])
	// 	fmt.Println(string(b[0:n]))

	// }

	//Write

	file, err:=os.OpenFile("/Users/chanchalverma/learning/core-packages-go/file-handling/temp.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err!= nil{
		fmt.Println(err)
	}
	defer file.Close()

	_,err= file.WriteString("Do you fine me useful?")
	if err!=nil{
		fmt.Println(err)
	}



}