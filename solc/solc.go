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
	PATH_JSON = "json"
	PATH_BIND = "bind"
)

func ExportBind(contracts map[string]*compiler.Contract, out string) error {
	var (
		abis  []string
		bins  []string
		types []string
		pkg   = "adapter"
		lang  = bind.LangGo
	)
	for n, contract := range contracts {
		var file = strings.Split(n, ":")[1]

		tmp, err := json.Marshal(contract.Info.AbiDefinition)
		if err != nil {
			return err
		}
		abis = []string{string(tmp)}
		bins = []string{contract.Code}
		types = []string{file}

		var code string
		if code, err = bind.Bind(types, abis, bins, pkg, lang); err != nil {
			return err
		}

		if err = utils.WriteFile(path.Join(out, PATH_BIND, file+".go"), []byte(code)); err != nil {
			return err
		}
	}
	return nil
}

func ExportJSON(contracts map[string]*compiler.Contract, out string) error {
	for n, contract := range contracts {
		var file = strings.Split(n, ":")[1] + ".json"
		var data, err = json.MarshalIndent(contract, "", "    ")
		if err != nil {
			return err
		}

		err = utils.WriteFile(path.Join(out, PATH_JSON, file), data)
		if err != nil {
			return err
		}
	}
	return nil
}

// Compile creates abi and bind of contract
func Compile(solc string, dir string, out string) (err error) {
	var contracts map[string]*compiler.Contract

	// Get solidity files from folder
	var sources []string
	if sources, err = utils.GetDirElems(dir); err != nil {
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
	if err = os.Chdir(".."); err != nil {
		return
	}

	if err = utils.RemoveDir(out); err != nil {
		return
	}

	if err = utils.CreateDirs(
		path.Join(out, PATH_JSON),
		path.Join(out, PATH_BIND),
	); err != nil {
		return
	}

	if err = ExportJSON(contracts, out); err != nil {
		return
	}

	if err = ExportBind(contracts, out); err != nil {
		return
	}

	return
}
