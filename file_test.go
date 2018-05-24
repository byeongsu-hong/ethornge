package ethornge

import (
	"os"
	"testing"
)

type TData struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

var tFile = &pathParam{"test/test.txt"}
var tCSV = &pathParam{"test/test.csv"}
var tJSON = &pathParam{"test/test.json"}

func TestCreateFile(t *testing.T) {
	createFile(tFile)

	if _, err := os.Stat(tFile.p); os.IsNotExist(err) {
		t.Errorf("Error : File not created")
	}
}

func TestWriteFile(t *testing.T) {

}

func TestWriteJSON(t *testing.T) {

}

func TestRemoveFile(t *testing.T) {

}

func TestReadFile(t *testing.T) {

}

func TestReadJSON(t *testing.T) {

}

func TestReadCSV(t *testing.T) {

}
