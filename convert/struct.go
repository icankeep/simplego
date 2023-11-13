package convert

import (
	"fmt"
	"github.com/icankeep/simplego/fmtx"
	"github.com/icankeep/simplego/setx"
	"regexp"
	"strings"
)

type StructInfo struct {
	Name       string
	StructBody string
	Fields     []*StructField
	Comment    string
}

type StructField struct {
	Name     string
	DataType string
	Tags     []*StructTag
	TagsStr  string
	Comment  string
}

type StructTag struct {
	TagType string
	Value   string
}

var (
	TestStruct       = "type Person struct {\n\tName     string `json:\"name\"`\n\tAge      int    `json:\"age\"`\n\tLocation string `json:\"location\"`\n}"
	StructPattern    = `\s*type\s+([^\s]+)\s+struct\s+{([^}]+)}`
	FieldLinePattern = "\\s+([^\\s]+)?\\s+([^\\s]+)([^\\n\\r]*)"
	TagsPattern      = "`(.+)`"
	TagPattern       = "([^\\s:]+):\"([^\"]+)\""
)

type StructParseHandler struct {
	Input      string `json:"input"`
	FmtInput   string
	Output     string
	AddTags    []string
	DeleteTags []string

	StructInfo
}

func (s *StructParseHandler) Parse(input string) (err error) {
	s.Input = input
	// 1. Format input
	s.FmtInput, err = fmtx.FormatGoCode(input)
	if err != nil {
		return err
	}

	// 2. Resolve input to struct
	if ok := s.parseStruct(); !ok {
		return fmt.Errorf("invalid struct")
	}
	if ok := s.parseFields(); !ok {
		return fmt.Errorf("invalid struct, cannot find field")
	}

	return nil
}

func (s *StructParseHandler) parseStruct() bool {
	reg := regexp.MustCompile(StructPattern)
	match := reg.FindStringSubmatch(s.FmtInput)
	if len(match) == 0 {
		return false
	}
	s.Name = match[1]
	s.StructBody = match[2]
	return true
}

func (s *StructParseHandler) parseFields() bool {
	reg := regexp.MustCompile(FieldLinePattern)

	match := reg.FindAllStringSubmatch(s.StructBody, -1)
	if len(match) == 0 {
		return false
	}
	fields := make([]*StructField, 0)
	for _, line := range match {
		tagsStr, tags := s.parseTags(line[3])
		fields = append(fields, &StructField{
			Name:     line[1],
			DataType: line[2],
			TagsStr:  tagsStr,
			Tags:     tags,
		})
	}
	s.Fields = fields
	return true
}

func (s *StructParseHandler) parseTags(str string) (tagsStr string, tags []*StructTag) {
	if len(str) == 0 {
		return "", tags
	}

	tagsReg := regexp.MustCompile(TagsPattern)
	tagReg := regexp.MustCompile(TagPattern)
	tagsMatch := tagsReg.FindStringSubmatch(str)

	if len(tagsMatch) == 0 {
		return "", tags
	}
	tagMatch := tagReg.FindAllStringSubmatch(tagsMatch[1], -1)
	for _, m := range tagMatch {
		tags = append(tags, &StructTag{
			TagType: m[1],
			Value:   m[2],
		})
	}
	return tagsMatch[1], tags
}

func (s *StructParseHandler) Handle(input string) (string, error) {
	if err := s.Parse(input); err != nil {
		return "", err
	}

	s.Output = s.FmtInput
	for _, field := range s.Fields {
		tagTypes := setx.Set[string]{}
		for _, tag := range field.Tags {
			tagTypes.Add(tag.TagType)
		}
		newTagsStr := field.TagsStr
		for _, tag := range s.AddTags {
			if tagTypes.Contains(tag) {
				continue
			}
			newTagsStr = fmt.Sprintf("%s %s:\"%s\"", newTagsStr, tag, GetTag(tag, field.Name))
		}
		for _, tag := range s.DeleteTags {
			if !tagTypes.Contains(tag) {
				continue
			}
			tagStr := fmt.Sprintf("%s:\"%s\"", tag, GetTag(tag, field.Name))
			newTagsStr = strings.Replace(newTagsStr, " "+tagStr, "", 1)
			newTagsStr = strings.Replace(newTagsStr, tagStr, "", 1)
		}
		newTagsStr = strings.TrimSpace(newTagsStr)
		s.Output = strings.Replace(s.Output, field.TagsStr, newTagsStr, 1)
	}
	return s.Output, nil
}

func GetTag(tagType, fieldName string) *StructTag {
	var value string

	switch tagType {
	case "json":
		value = UnderscoreToLowerCamelCase(fieldName)
	case "gorm":
		value = fieldName
	case "xml":
		value = fieldName
	case "yaml":
		value = UnderscoreToUpperCamelCase(fieldName)
	default:
		value = UnderscoreToUpperCamelCase(fieldName)
	}

	return &StructTag{
		TagType: tagType,
		Value:   value,
	}
}
