package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsEmpty(t *testing.T) {
	a := assert.New(t)

	a.True(IsEmpty(""))
	a.False(IsEmpty("TEST"))
	a.False(IsEmpty(" "))
	a.False(IsEmpty("\t"))
	a.False(IsEmpty("\n"))
	a.False(IsEmpty("꯱"))
}

func Test_IsNotEmpty(t *testing.T) {
	a := assert.New(t)

	a.False(IsNotEmpty(""))
	a.True(IsNotEmpty("TEST"))
	a.True(IsNotEmpty(" "))
	a.True(IsNotEmpty("\t"))
	a.True(IsNotEmpty("\n"))
	a.True(IsNotEmpty("꯱"))
}

func Test_IsBlank(t *testing.T) {
	a := assert.New(t)

	a.True(IsBlank(""))
	a.True(IsBlank(" "))
	a.True(IsBlank("\t"))
	a.True(IsBlank("\n"))
	a.True(IsBlank("\n  \t"))
	a.True(IsBlank("\n  "))
	a.True(IsBlank("  \t  "))
	a.True(IsBlank("     "))
	a.False(IsBlank("  s   "))
	a.False(IsBlank("xxx"))
}

func Test_IsNotBlank(t *testing.T) {
	a := assert.New(t)

	a.False(IsNotBlank(""))
	a.False(IsNotBlank(" "))
	a.False(IsNotBlank("\t"))
	a.False(IsNotBlank("\n"))
	a.False(IsNotBlank("\n  \t"))
	a.False(IsNotBlank("\n  "))
	a.False(IsNotBlank("  \t  "))
	a.False(IsNotBlank("     "))
	a.True(IsNotBlank("  s   "))
	a.True(IsNotBlank("xxx"))
}

func Test_IsNumeric(t *testing.T) {
	a := assert.New(t)

	a.True(IsNumeric("0"))
	a.True(IsNumeric("2"))
	a.True(IsNumeric("9"))
	a.False(IsNumeric("X"))
	a.False(IsNumeric("1x"))
	a.False(IsNumeric("0x1234"))
	a.False(IsNumeric(" "))
}
