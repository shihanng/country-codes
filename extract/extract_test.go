package extract

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdata = "testdata/"

func TestExtractAlpha2Code(t *testing.T) {
	f, err := os.Open(testdata + "alpha2_code.html")
	require.NoError(t, err)
	defer f.Close()

	got, err := ExtractAlpha2Code(f)
	assert.NoError(t, err)

	gotBlob, err := json.MarshalIndent(got, "", "  ")
	assert.NoError(t, err)

	g := goldie.New(t, goldie.WithNameSuffix(".golden.json"))
	g.Assert(t, "alpha2_code", gotBlob)
}

func TestExtractDetail(t *testing.T) {
	f, err := os.Open(testdata + "detail.html")
	require.NoError(t, err)
	defer f.Close()

	got, err := ExtractDetail(f)
	assert.NoError(t, err)

	gotBlob, err := json.MarshalIndent(got, "", "  ")
	assert.NoError(t, err)

	g := goldie.New(t, goldie.WithNameSuffix(".golden.json"))
	g.Assert(t, "extract_detail", gotBlob)
}
