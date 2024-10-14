package main

import (
	// "fmt"
	// "log"
	"github.com/sirupsen/logrus"
	// "os"
)

/*
log
glog
logrus
*/

func main(){

	// log.Println("Hello Chanchal")

	// file, err:=os.OpenFile("/Users/chanchalverma/learning/core-packages-go/logging-package/app.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	
	// if err!=nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// log.SetOutput(file)
	// log.Println("Hello Chanchal....!")
	// logrus.SetOutput(file)
	logrus.Fatal("Hello Chanchal....!")
	// logrus.SetLevel(logrus.Level())
	
	
}

/*
Log level
-Trace
-Debug
-Info
-Warn
-Error
-Fatal
-Panic

*/