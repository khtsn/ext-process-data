package main

import (
	"flag"
	"fmt"
	"github.com/khtsn/ext-process-data/process"
	"github.com/khtsn/ext-process-data/utils"
)

func main() {
	//1. parse inputs
	//1.1. get file path
	filePathPtr := flag.String("file", "testdata/test.csv", "input csv file")
	//1.2. custom result name if any
	outputPtr := flag.String("output", "result", "name the result file name on csv format")
	//1.3. parse ptr
	flag.Parse()
	filePath := *filePathPtr
	output := *outputPtr

	//1.5. show configs
	fmt.Println("input file path:", filePath)
	fmt.Println("output csv result file name:", output)

	//2. create output temp file
	tmpFile := utils.CreateFile(output + "_tmp.csv")
	defer tmpFile.Close()

	//3. process file path provided
	proc := process.NewProcess(tmpFile)
	err := proc.ProcessFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("ProcessFile error: %s", err.Error()))
	}

	//4. output to final result
	err = proc.Export(output + ".csv")
	if err != nil {
		panic(fmt.Sprintf("Export error: %s", err.Error()))
	}
}
