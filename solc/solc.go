package solc

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/frostornge/ethornge/utils"
)

const (
	outPath = "out"
	adtPath = "adapter"
)

func exportBind(contracts map[string]*compiler.Contract, out string) (err error) {
	var (
		abis  []string
		bins  []string
		types []string
		pkg   = "adapter"
		lang  = bind.LangGo
	)
	for n := range contracts {
		var file = strings.Split(n, ":")[1]
		var abiPath = path.Join(out, outPath, file+".json")
		var binPath = path.Join(out, outPath, file+".bin")
		var bindPath = path.Join(out, adtPath, file+".go")

		var data []byte

		if data, err = utils.ReadFile(abiPath); err != nil {
			return
		}
		abis = []string{string(data)}

		if data, err = utils.ReadFile(binPath); err != nil {
			return
		}
		bins = []string{string(data)}
		types = []string{file}

		var code string
		if code, err = bind.Bind(types, abis, bins, pkg, lang); err != nil {
			return
		}

		if err = utils.WriteFile(bindPath, []byte(code)); err != nil {
			return
		}
	}
	return
}

func exportBIN(contracts map[string]*compiler.Contract, out string) (err error) {
	for n, contract := range contracts {
		var file = strings.Split(n, ":")[1] + ".bin"

		if err = utils.WriteFile(
			path.Join(out, outPath, file),
			[]byte(contract.Code),
		); err != nil {
			return
		}
	}
	return
}

func exportABI(contracts map[string]*compiler.Contract, out string) (err error) {
	for n, contract := range contracts {
		var data []byte
		var file = strings.Split(n, ":")[1] + ".json"

		if data, err = json.MarshalIndent(
			contract.Info.AbiDefinition,
			"",
			"    ",
		); err != nil {
			return
		}

		if err = utils.WriteFile(
			path.Join(out, outPath, file),
			data,
		); err != nil {
			return
		}
	}
	return
}

// Compile creates abi and bind of contract
func Compile(solc string, dir string, out string) (err error) {
	var contracts map[string]*compiler.Contract

	// Create Directory
	if err = utils.CreateDir(path.Join(out, outPath)); err != nil {
		return
	}

	if err = utils.CreateDir(path.Join(out, adtPath)); err != nil {
		return
	}

	// Get solidity files from folder
	var sources []string
	if sources, err = utils.GetDirElems(dir); err != nil {
		return
	}

	// Get current working path
	var origin string
	if origin, err = os.Getwd(); err != nil {
		return
	}

	// Change working directory which contains solidity files
	if err = os.Chdir(dir); err != nil {
		return
	}

	if contracts, err = compiler.CompileSolidity(solc, sources...); err != nil {
		return
	}

	// Restore working path
	if err = os.Chdir(origin); err != nil {
		return
	}

	if err = exportABI(contracts, out); err != nil {
		return
	}

	if err = exportBIN(contracts, out); err != nil {
		return
	}

	if err = exportBind(contracts, out); err != nil {
		return
	}

	return
}
