package test

import (
	"os"
)

//MyLogFile fir storing log entries
type MyLogFile struct {
	name string
	file *os.File
}

//LogFile ???
var LogFile MyLogFile

//Init to set up the file
func init() {
	LogFile.name = "Test.log"
	f, err := os.Create(LogFile.name)
	LogFile.file = f
	if err != nil {
		println(LogFile.name, err)
	}
}

func (f *MyLogFile) Write(s string) {
	_, err := f.file.Write([]byte(s))
	if err != nil {
		println(f.name, err)
	}
}

//Close ???
func (f *MyLogFile) Close() {
	err := f.file.Close()
	if err != nil {
		println(f.name, err)
	}
}
