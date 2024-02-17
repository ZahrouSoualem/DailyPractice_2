package main

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	FatalLog   *log.Logger
)

func init() {
	file, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLog = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLog.Println("Opening the application...")
	InfoLog.Println("Something has occurred...")
	WarningLog.Println("WARNING!!!..")
	ErrorLog.Println("Some error has occurred...")
	/* log.Fatal("Critical error. Exiting.") */
	FatalLog.Fatal("Some Fatal error has occurred...")
	status := 404
	log.Fatalf("Fatal error. HTTP status: %d", status)
	fmt.Println("Hello world")
}
