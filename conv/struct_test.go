package conv

import (
	"fmt"
	"github.com/icankeep/simplego/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructConvert_1(t *testing.T) {
	s := &StructParseHandler{
		AddTags:    []string{"xml", "json"},
		DeleteTags: []string{"json"},
	}
	s.testStructConvert(t, "1", false)
}

func TestStructConvert_2(t *testing.T) {
	s := &StructParseHandler{
		AddTags: []string{"xml", "json"},
	}
	s.testStructConvert(t, "2", false)
}

func TestStructConvert_3(t *testing.T) {
	s := &StructParseHandler{
		AddTags: []string{"xml", "json"},
	}
	s.testStructConvert(t, "3", false)
}

func TestStructConvert_4(t *testing.T) {
	s := &StructParseHandler{
		AddTags: []string{"yaml"},
	}
	s.testStructConvert(t, "4", false)
}

func (s *StructParseHandler) testStructConvert(t *testing.T, testKey string, print bool) {
	a := assert.New(t)
	output, err := s.Handle(utils.MustReadStringFromFile(fmt.Sprintf("../utest/data/TestStructConvert_%v_input.txt", testKey)))

	a.NoError(err, fmt.Sprintf("%+v", err))

	if print {
		fmt.Println(output)
	}
	a.Equal(utils.MustReadStringFromFile(fmt.Sprintf("../utest/data/TestStructConvert_%v_output.txt", testKey)), output)
}
