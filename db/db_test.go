package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(Memory)
	require.NoError(t, err)

	_ = db
}
