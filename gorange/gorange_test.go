package gorange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLaunch(t *testing.T) {
	n, err := Launch(DefaultLocalConfig(nil, 0))
	assert.NoError(t, err)

	_, err = n.WsProvider(context.Background())
	assert.NoError(t, err)

	n.Stop()
}
