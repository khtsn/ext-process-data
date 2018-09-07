# ext-process-data
External process data in file

## Environment
* Go 1.11
* ArchLinux / macOS Sierra

## Usage
Make sure $GOPATH is defined
* go get github.com/khtsn/ext-process-data
* ./ext-process-data -file=testdata/test.csv -output=result.csv

## Future improvements
* Multiple tmp files for processing, with index file for lookup position
* Using goroutines for multiple processing