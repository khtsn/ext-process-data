package process

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/khtsn/ext-process-data/utils"
	"os"
	"strings"
)

const ErrorPattern = "error %v on line %d"

var lineNumber uint = 0 //for error checking

type Process struct {
	File *os.File
}

// NewProcess init file process object with temp file
func NewProcess(f *os.File) *Process {
	return &Process{
		File: f,
	}
}

// ProcessFile read file and pass line data to line parser
func (p *Process) ProcessFile(filePath string) error {
	//open file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	//close file after get things done
	defer file.Close()

	//scan line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber++
		if lineNumber == 1 { //first line header
			continue
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "" { //empty line
			continue
		}

		phoneDetail, err := utils.ParseLine(line)
		if err != nil {
			return errors.New(fmt.Sprintf(ErrorPattern, err, lineNumber))
		}

		err = p.ProcessPhonePlan(&phoneDetail)
		if err != nil {
			return errors.New(fmt.Sprintf(ErrorPattern, err, lineNumber))
		}
	}

	//reach to the end of file or got error
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Export export tmp file to final result
func (p *Process) Export(output string) error {
	//retro cursor to beginning of file
	p.File.Seek(0, 0)
	//read all data from tmp file
	reader := csv.NewReader(p.File)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	//create new file
	f := utils.CreateFile(output)
	//use csv writer to file
	w := csv.NewWriter(f)
	//write header first
	w.Write(strings.Split(utils.HeaderOutputCSV, ","))
	//write just record[0] and record[1]
	for _, record := range records {
		w.Write([]string{record[0], record[1]})
	}
	w.Flush() //write all

	//clean tmp file
	os.Remove(p.File.Name())

	return nil
}
