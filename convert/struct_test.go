package convert

import "testing"

func TestStructConvert(t *testing.T) {
	s := &StructParseHandler{}
	s.Parse("type StructParseHandler struct {\n\tInput      string `json:\"input\"`\n\tFmtInput   string\n\tOutput     string\n\tAddTags    []string\n\tDeleteTags []string\n\n\tStructInfo\n}")
}
