package testornge

import (
	"strings"

	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func parse(str string) (keys []string) {
	for _, s := range strings.Split(str, "\n") {
		if res, _ := regexp.MatchString(`\(\d\)`, s); res {
			var valid = strings.Split(s, " ")[1]
			if !common.IsHexAddress(valid) {
				keys = append(keys, valid)
			}
		}
	}
	return
}
