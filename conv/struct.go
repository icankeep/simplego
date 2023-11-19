package conv

import (
	"fmt"
	"github.com/icankeep/simplego/fmtx"
	"github.com/icankeep/simplego/setx"
	"github.com/icankeep/simplego/slicex"
	"github.com/icankeep/simplego/utils"
	"regexp"
	"strings"
)

type StructInfo struct {
	Name       string
	StructBody string
	Fields     []*StructField
	Comment    string
}

func (s *StructInfo) String() string {
	fieldLines := make([]string, 0)
	for _, field := range s.Fields {
		tagStrs := make([]string, 0)
		for _, tag := range field.Tags {
			tagStrs = append(tagStrs, fmt.Sprintf(StructTagTemplate, tag.TagType, tag.Value))
		}
		comment := ""
		if field.Comment != "" {
			comment = fmt.Sprintf("// %s", field.Comment)
		}
		fieldLine := fmt.Sprintf(StructLineTemplate, field.Name, field.DataType, strings.Join(tagStrs, " "), comment)
		fieldLines = append(fieldLines, fieldLine)
	}
	comment := ""
	if s.Comment != "" {
		comment = fmt.Sprintf("// %s %s", s.Name, s.Comment)
	}
	return fmt.Sprintf(StructTemplate, comment, s.Name, strings.Join(fieldLines, ""))
}

type StructField struct {
	RealName string
	Name     string
	DataType string
	Tags     []*StructTag
	TagsStr  string
	LineStr  string
	Comment  string
}

type StructTag struct {
	TagType string
	Value   string
}

var (
	StructPattern       = `(?s)\s*type\s+([^\s]+)\s+struct\s+{(.*?)\n}`
	FieldLinePattern    = "\\s+(([^\\s]+)?\\s+([^\\n/`]+)([^\\n\\r]*))"
	FieldCommentPattern = "(//[^\\n\\r]+)"
	TagsPattern         = "`(.+)`"
	TagPattern          = "([^\\s:]+):\"([^\"]+)\""
)

type StructParseHandler struct {
	Input      string
	FmtInput   string
	Output     string
	AddTags    []string
	DeleteTags []string

	StructInfo
}

func RemoveTags(input string, tags []string) (string, error) {
	h := &StructParseHandler{
		DeleteTags: tags,
	}
	return h.handle(input)
}

func AddTags(input string, tags []string) (string, error) {
	h := &StructParseHandler{
		AddTags: tags,
	}
	return h.handle(input)
}

func (s *StructParseHandler) parse(input string) (err error) {
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
	fieldCommentReg := regexp.MustCompile(FieldCommentPattern)

	match := reg.FindAllStringSubmatch(s.StructBody, -1)
	if len(match) == 0 {
		return false
	}
	fields := make([]*StructField, 0)
	for _, line := range match {
		if strings.HasPrefix(strings.TrimSpace(line[0]), "//") {
			continue
		}
		commentMatch := fieldCommentReg.FindStringSubmatch(line[0])
		tagsStr, tags := s.parseTags(line[4])
		fields = append(fields, &StructField{
			Name:     line[2],
			DataType: strings.TrimSpace(line[3]),
			TagsStr:  tagsStr,
			Tags:     tags,
			LineStr:  line[1],
			Comment:  utils.SafeIndexValueOrDefault(commentMatch, 1),
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

func (s *StructParseHandler) handle(input string) (string, error) {
	if err := s.parse(input); err != nil {
		return "", fmt.Errorf("parse struct error, %v", err)
	}

	s.addOrRemoveTags()

	return fmtx.FormatGoCode(s.Output)
}

func (s *StructParseHandler) addOrRemoveTags() {
	s.Output = s.FmtInput
	for _, field := range s.Fields {
		if len(field.Name) == 0 {
			continue
		}
		tagTypes := setx.NewSet[string]()
		for _, tag := range field.Tags {
			tagTypes.Add(tag.TagType)
		}
		newTagsStr := field.TagsStr
		for _, tag := range s.AddTags {
			if tagTypes.Contains(tag) {
				continue
			}
			newTagsStr = fmt.Sprintf("%s %s:\"%s\"", newTagsStr, tag, getTagValue(tag, field.Name))
		}
		for _, tag := range s.DeleteTags {
			tagStr := fmt.Sprintf("%s:\"%s\"", tag, getTagValue(tag, field.Name))
			newTagsStr = strings.Replace(newTagsStr, " "+tagStr, "", 1)
			newTagsStr = strings.Replace(newTagsStr, tagStr, "", 1)
		}
		newTagsStr = strings.TrimSpace(newTagsStr)
		newLine := fmt.Sprintf("%s %s", field.Name, field.DataType)
		if len(newTagsStr) != 0 {
			newLine += " `" + newTagsStr + "`"
		}
		if len(field.Comment) != 0 {
			newLine += " " + field.Comment
		}

		s.Output = strings.Replace(s.Output, field.LineStr, newLine, 1)
	}
}

func getTagValue(tagType, fieldName string) string {
	return getTag(tagType, fieldName, nil).Value
}

func getTag(tagType, fieldName string, originNameTags []string) *StructTag {

	var value string
	switch tagType {
	case "json":
		value = UnderscoreToUpperCamelCase(fieldName)
	case "gorm":
		value = utils.If[string](slicex.Contains(originNameTags, tagType), fieldName, CamelCaseToUnderscore(fieldName))
		value = "column:" + value
	case "xml":
		value = CamelCaseToUnderscore(fieldName)
	case "yaml":
		value = UnderscoreToUpperCamelCase(fieldName)
	default:
		value = UnderscoreToUpperCamelCase(fieldName)
	}
	value = ProcessIDForFieldName(value)
	return &StructTag{
		TagType: tagType,
		Value:   value,
	}
}
