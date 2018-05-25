package ethornge

import (
	"testing"
)

func TestNewLedgerProvider(t *testing.T) {
	var err error

	if err = NewLedgerProvider(); err != nil {
		t.Error("Error : ", err)
	}
}
