package utils

import (
	"os"
)

// CreateFile create a file in current root
func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

// WriteToFile write data at cursor
func WriteToFile(file *os.File, content string, cursor int64) error {
	_, err := file.WriteAt([]byte(content), cursor)
	if err != nil {
		return err
	}
	return nil
}

// AppendToFile write string data to the end of file
func AppendToFile(file *os.File, content string) error {
	file.Seek(0, 2)
	_, err := file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
