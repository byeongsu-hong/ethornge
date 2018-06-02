package solc

import (
	"os"
	"path"
	"strings"
	"testing"

	"../utils"
)

var TestContractPath = "../test/contract"
var TestBuildPath = "../test/build"

func TestCompile(t *testing.T) {
	utils.RemoveDir("test/build")

	if err := Compile(
		"solc",
		TestContractPath,
		TestBuildPath,
	); err != nil {
		t.Error("Error : ", err)
	}

	// Get solidity files from folder
	var err error
	var sources []string
	if sources, err = utils.GetDirElems(TestContractPath); err != nil {
		return
	}

	for _, source := range sources {
		var name = strings.TrimSuffix(source, path.Ext(source))
		var abiPath = path.Join(TestBuildPath, "out", name+".json")
		var binPath = path.Join(TestBuildPath, "out", name+".bin")
		var bindPath = path.Join(TestBuildPath, "adapter", name+".go")

		if _, err = os.Stat(abiPath); os.IsNotExist(err) {
			t.Error("Error : ABI not found")
		}
		if _, err = os.Stat(binPath); os.IsNotExist(err) {
			t.Error("Error : BIN not found")
		}
		if _, err = os.Stat(bindPath); os.IsNotExist(err) {
			t.Error("Error : Bind not found")
		}
	}
}
