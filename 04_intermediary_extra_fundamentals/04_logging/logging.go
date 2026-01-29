package main

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	// general log
	log.Printf("This is a log messgae\n")

	// log with prefix
	log.SetPrefix("INFO: ")
	log.Printf("This is a log messgae\n")

	// log with flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("This is a log messgae\n")

	// log with prefix and flags
	log.SetPrefix("DEBUG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("This is a log messgae\n")

	// using custom loggers
	infoLogger.Println("This is an info message")
	debugLogger.Println("This is a debug message")
	errorLogger.Println("This is an error message")
	warnLogger.Println("This is a warn message")


	// create a new file for logs with append mode and read/write permissions
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}

	// defer file closing
	defer file.Close()

	// set output to file
	infoLogger.SetOutput(file)

	// log to file
	infoLogger.Println("This is an info message to file")
}
