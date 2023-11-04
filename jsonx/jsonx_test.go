package jsonx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatJSON(t *testing.T) {
	a := assert.New(t)

	content := `["123", "456"]`
	expectContent := `[
    "123",
    "456"
]`
	value, err := FormatJSON(content, 4)
	a.NoError(err)
	a.Equal(value, expectContent)
}
