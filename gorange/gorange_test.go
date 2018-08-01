package gorange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLaunch(t *testing.T) {
	n, err := Launch(Config{
		NetworkId: 2470,
		WSHost:    "127.0.0.1",
		WSPort:    2471,
		HTTPHost:  "127.0.0.1",
		HTTPPort:  2472,
		Accounts:  20,
		Balances:  100,
	})
	assert.NoError(t, err)

	_, err = n.WsProvider(context.Background())
	assert.NoError(t, err)

	n.Stop()
}
