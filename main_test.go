package main

import (
	"fmt"
	"github.com/khtsn/ext-process-data/process"
	"github.com/khtsn/ext-process-data/utils"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := `PHONE_NUMBER,ACTIVATION_DATE,DEACTIVATION_DATE
0987000001,2016-03-01,2016-05-01
0987000002,2016-02-01,2016-03-01
0987000001,2016-01-01,2016-03-01
0987000001,2016-12-01,
0987000002,2016-03-01,2016-05-01
0987000003,2016-01-01,2016-01-10
0987000001,2016-11-01,2016-12-01
0987000002,2016-05-01,
0987000001,2016-09-01,2016-11-01
0987000001,2016-07-01,2016-09-01
0987000001,2016-06-01,2016-07-01`

	output := `PHONE_NUMBER,REAL_ACTIVATION_DATE
0987000001,2016-06-01
0987000002,2016-02-01
0987000003,2016-01-01`

	//generate input data
	inputFileName := "main_test_multi.csv"
	inputTmpFileName := "main_test_multi_tmp.csv"
	f := utils.CreateFile(inputFileName)
	defer f.Close()
	utils.AppendToFile(f, input)

	//temp file
	fTmp := utils.CreateFile(inputTmpFileName)
	defer fTmp.Close()

	//process data
	p := process.NewProcess(fTmp)
	p.ProcessFile(inputFileName)

	//export result
	outputFileName := "main_test_multi_result.csv"
	p.Export(outputFileName)

	//read output and verify
	content, _ := ioutil.ReadFile(outputFileName)
	if strings.TrimSpace(string(content)) != output {
		fmt.Println(strings.TrimSpace(string(content)))
		t.Fail()
	}

	//clean files
	os.Remove(inputFileName)
	os.Remove(inputTmpFileName)
	os.Remove(outputFileName)
}
