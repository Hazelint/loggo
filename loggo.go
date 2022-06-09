package loggo

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	@title: Function Fatal
	@description: Log a fatal error
	@params: error
*/
func Fatal(data any) {
	// Get dir of log file
	dir := checkFile()

	// Open the log file
	file, err1 := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file.Close()
	// Get timestamp for log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// Write new log to file
	if _, err := file.WriteString("[" + timestamp + "] Fatal error: " + fmt.Sprint(data) + "\n"); err != nil {
		log.Fatal(err)
	}
}

/*
	@title: Function Fatal
	@description: Log a fatal error
	@params: error
*/
func Info(data any) {
	// Get dir of log file
	dir := checkFile()

	// Open the log file
	file, err := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get timestamp for log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// Write new log to file
	if _, err := file.WriteString("[" + timestamp + "] Info: " + fmt.Sprint(data) + "\n"); err != nil {
		log.Fatal(err)
	}
}

func Greet() (string, any) {
	return "Hello", nil
}

/*
	@title: Function checkFile
	@description: Check if log file exists otherwise create a new one
	@return: string
*/
func checkFile() string {
	// Declare variables
	var createFile bool

	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		createFile = true
	}

	// Create a new file if file does not exists
	if createFile {
		file, err := os.Create(dir + "/go.log")
		if err != nil {
			log.Fatal(err)
		}
		// Write first line
		file.WriteString("Start log:\n")
		// Return directory
		return dir + "/go.log"
	}
	// Return directory if file exists because it is expected
	return dir + "/go.log"

}
