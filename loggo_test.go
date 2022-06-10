package loggo

import (
	"os"
	"testing"
)

func TestFatal(t *testing.T) {
	errorTest := "test fatal"
	Fatal(errorTest)
	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Directory not found")
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		t.Fatal("File not created")
	}

	// Remove file for next test
	err3 := os.Remove(dir + "/go.log")
	if err3 != nil {
		t.Fatal("Test file not found")
	}
}

func TestError(t *testing.T) {
	errorTest := "test error"
	Error(errorTest)
	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Directory not found")
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		t.Fatal("File not created")
	}

	// Remove file for next test
	err3 := os.Remove(dir + "/go.log")
	if err3 != nil {
		t.Fatal("Test file not found")
	}
}

func TestWarn(t *testing.T) {
	errorTest := "test warn"
	Warn(errorTest)
	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Directory not found")
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		t.Fatal("File not created")
	}

	// Remove file for next test
	err3 := os.Remove(dir + "/go.log")
	if err3 != nil {
		t.Fatal("Test file not found")
	}
}

func TestInfo(t *testing.T) {
	infoTest := "test info"
	Info(infoTest)
	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Directory not found")
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		t.Fatal("File not created")
	}

	// Remove file for next test
	err3 := os.Remove(dir + "/go.log")
	if err3 != nil {
		t.Fatal("Test file not found")
	}
}

func TestDebug(t *testing.T) {
	errorTest := "test debug"
	Debug(errorTest)
	// Get the directory where the log file is expected
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Directory not found")
	}

	// Check if log file is at expected location
	_, err2 := os.Stat(dir + "/go.log")
	if os.IsNotExist(err2) {
		// Set createFile to true if file does not exists
		t.Fatal("File not created")
	}

	// Remove file for next test
	err3 := os.Remove(dir + "/go.log")
	if err3 != nil {
		t.Fatal("Test file not found")
	}
}
