package ethornge

import (
	"os"
	"strings"

	"encoding/json"

	"path"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
)

const (
	outPath = "out"
	adtPath = "adapter"
)

func exportBind(sources []string, out string) (err error) {
	var (
		abis  []string
		bins  []string
		types []string
		pkg   = "adapter"
		lang  = bind.LangGo
	)
	for _, source := range sources {
		var name = strings.TrimSuffix(source, path.Ext(source))
		var abiPath = path.Join(out, outPath, name+".json")
		var binPath = path.Join(out, outPath, name+".bin")
		var bindPath = path.Join(out, adtPath, name+".go")

		var data []byte

		if data, err = readFile(abiPath); err != nil {
			return
		}
		abis = []string{string(data)}

		if data, err = readFile(binPath); err != nil {
			return
		}
		bins = []string{string(data)}
		types = []string{name}

		var code string
		if code, err = bind.Bind(types, abis, bins, pkg, lang); err != nil {
			return
		}

		if err = writeFile(bindPath, []byte(code)); err != nil {
			return
		}
	}
	return
}

func exportBIN(sources []string, contracts map[string]*compiler.Contract, out string) (err error) {
	for _, source := range sources {
		var name = strings.TrimSuffix(source, path.Ext(source))

		if err = writeFile(
			path.Join(out, outPath, name+".bin"),
			[]byte(contracts[name+".sol:"+name].Code),
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

		if err = writeFile(
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
	if err = createDir(path.Join(out, outPath)); err != nil {
		return
	}

	if err = createDir(path.Join(out, adtPath)); err != nil {
		return
	}

	// Get solidity files from folder
	var sources []string
	if sources, err = getDirElems(dir); err != nil {
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

	if err = exportBIN(sources, contracts, out); err != nil {
		return
	}

	if err = exportBind(sources, out); err != nil {
		return
	}

	return
}
