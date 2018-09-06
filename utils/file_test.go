package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFileSuccess(t *testing.T) {
	//1. create a file
	fileName := "test_create_file.csv"
	f := CreateFile(fileName)
	defer f.Close()
	//2. verify if the file exist
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatal("cannot create the file")
	}
	//3. clean test files
	os.Remove(fileName)
}

func TestAppendToFile(t *testing.T) {
	line := "hello world\nthis is line 2"
	//1. create a file
	fileName := "test_append_file.txt"
	f := CreateFile(fileName)
	defer f.Close()

	//2. append data
	AppendToFile(f, line)

	//3. verify
	reader, _ := ioutil.ReadFile(fileName)

	if line != string(reader) {
		t.Fail()
	}

	//clean test files
	os.Remove(fileName)
}
