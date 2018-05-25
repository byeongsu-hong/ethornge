package ethornge

import (
	"fmt"
	"os"
	"testing"
)

type TestStruct struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

var TestData = TestStruct{
	A: "A",
	B: "B",
	C: "C",
	D: "D",
}
var TestFile = "test/test.txt"
var TestCSV = "test/test.csv"

func TestCreateFile(t *testing.T) {
	var err error

	if err = removeFile(TestFile); err != nil {
		t.Error("Error : ", err)
	}
	if err = createFile(TestFile); err != nil {
		t.Error("Error : ", err)
	}
	if _, err = os.Stat(TestFile); os.IsNotExist(err) {
		t.Error("Error : File not created")
	}
}

func TestReadWriteRemoveFile(t *testing.T) {
	var err error
	var data []byte

	if err = removeFile(TestFile); err != nil {
		t.Error("Error : ", err)
	}
	if err = writeFile(TestFile, []byte(fmt.Sprint(TestData))); err != nil {
		t.Error("Error : ", err)
	}
	if data, err = readFile(TestFile); err != nil {
		t.Error("Error : ", err)
	}
	if string(data) != fmt.Sprint(TestData) {
		t.Error("Error : Write wrong data on file")
	}
}

func TestReadCSV(t *testing.T) {
	var vs, err = readCSV(TestCSV, TestStruct{})
	if err != nil {
		t.Error("Error : ", err)
	}

	var res []TestStruct
	for _, v := range vs {
		if s, ok := v.(TestStruct); ok {
			res = append(res, s)
		}
	}

	if len(res) == 3 {
		t.Error("Error : Read error")
	}
}
