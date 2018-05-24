package ethornge

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
)

type contracts map[string]*compiler.Contract

func (c contracts) getAbis() (abi []string) {
	var exclude = make(map[string]bool)
	for name, contract := range c {
		if exclude[strings.ToLower(name)] {
			continue
		}

		var tmp, err = json.Marshal(contract.Info.AbiDefinition)
		if err != nil {
			log.Fatal(err)
		}

		abi = append(abi, string(tmp))
	}
	return
}

func (c contracts) getBins() (bin []string) {
	var exclude = make(map[string]bool)
	for name, contract := range c {
		if exclude[strings.ToLower(name)] {
			continue
		}

		bin = append(bin, contract.Code)
	}
	return
}

func (c contracts) getTypes() (types []string) {
	var exclude = make(map[string]bool)
	for name := range c {
		if exclude[strings.ToLower(name)] {
			continue
		}

		var nameParts = strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])
	}
	return
}

func (c contracts) getAll() (abi, bin, types []string) {
	var exclude = make(map[string]bool)
	for name, contract := range c {
		if exclude[strings.ToLower(name)] {
			continue
		}

		var tmp, err = json.Marshal(contract.Info.AbiDefinition)
		if err != nil {
			log.Fatal(err)
		}

		abi = append(abi, string(tmp))
		bin = append(bin, contract.Code)

		var nameParts = strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])
	}
	return
}

func contractToAbi(solc string, contract *pathParam, out *pathParam) error {
	if solc == "" || contract == nil || out == nil {
		return fmt.Errorf("Invalid arguments")
	}

	var cs, err = compiler.CompileSolidity(solc, contract.p)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, abi := range contracts(cs).getAbis() {
		builder.WriteString(abi)
	}

	writeFile(out, []byte(builder.String()))
	return nil
}

func contractToBin(solc string, contract *pathParam, out *pathParam) error {
	if solc == "" || contract == nil || out == nil {
		return fmt.Errorf("Invalid arguments")
	}

	var cs, err = compiler.CompileSolidity(solc, contract.p)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, bin := range contracts(cs).getBins() {
		builder.WriteString(bin)
	}

	writeFile(out, []byte(builder.String()))
	return nil
}

func abiToBind(abi *pathParam, typ string, pkg string, out *pathParam) error {
	if abi == nil || out == nil || pkg == "" {
		return fmt.Errorf("Invalid arguments")
	}

	if typ == "" {
		typ = pkg
	}

	var abis = []string{string(readFile(abi))}
	var bins []string
	var types = []string{typ}

	var code, err = bind.Bind(types, abis, bins, pkg, bind.LangGo)
	if err != nil {
		return err
	}
	writeFile(out, []byte(code))
	return nil
}

func contractToBind(solc string, contract *pathParam, typ string, pkg string, out *pathParam) error {
	if solc == "" || contract == nil || out == nil || pkg == "" {
		return fmt.Errorf("Invalid arguments")
	}

	if typ == "" {
		typ = pkg
	}

	var err error
	var cs contracts

	cs, err = compiler.CompileSolidity(solc, contract.p)
	if err != nil {
		return err
	}

	var abis, bins, _ = cs.getAll()
	var code string

	code, err = bind.Bind([]string{typ}, abis, bins, pkg, bind.LangGo)
	if err != nil {
		return err
	}
	writeFile(out, []byte(code))
	return nil
}
