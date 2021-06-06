package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(context.Background(), Memory)
	require.NoError(t, err)

	_ = db
}
