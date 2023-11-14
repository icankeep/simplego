package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructConvert(t *testing.T) {
	a := assert.New(t)
	s := &StructParseHandler{
		AddTags:    []string{"xml", "json"},
		DeleteTags: []string{"json"},
	}
	_, err := s.Handle("type StructParseHandler struct {\n\tInput      string `json:\"input\"`     // Required ...\n\tFmtInput   string `json:\"fmt_input\"` // Required ssss ...\n\tOutput     string\n\tAddTags    []string\n\tDeleteTags []string\n\n\tStructInfo\n}")
	a.NoError(err)

	//fmt.Println(output)
}
