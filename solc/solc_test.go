package solc

import (
	"testing"

	"github.com/frostornge/ethornge/utils"
)

var TestContractPath = "../test/contract"
var TestBuildPath = "../test/build"

func TestCompile(t *testing.T) {
	utils.RemoveDir(TestBuildPath)

	if err := Compile(
		"solc",
		TestContractPath,
		TestBuildPath,
	); err != nil {
		t.Error("Error : ", err)
	}
}
