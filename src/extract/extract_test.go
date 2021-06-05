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
		Alpha2Code:         "BR",
		ShortName:          "BRAZIL",
		ShortNameLowerCase: "Brazil",
		FullName:           "the Federative Republic of Brazil",
		Alpha3Code:         "BRA",
		NumericCode:        "076",
		Independent:        "Yes",
		TerritoryName:      "Fernando de Noronha Island, Martim Vaz Islands, Trindade Island",
		Status:             "Officially assigned",
	}

	got, err := ExtractDetail(f)
	assert.NoError(t, err)
	assert.Equal(t, &expected, got)
}
