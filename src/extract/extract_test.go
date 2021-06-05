package extract

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdata = "testdata/"

func TestExtractDetail(t *testing.T) {
	f, err := os.Open(testdata + "detail.html")
	require.NoError(t, err)
	defer f.Close()

	expected := Detail{
		Alpha2Code: "BR",
		ShortName:  "BRAZIL",
	}

	got, err := ExtractDetail(f)
	assert.NoError(t, err)
	assert.Equal(t, &expected, got)
}
