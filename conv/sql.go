package conv

import (
	"fmt"
	"github.com/icankeep/simplego/fmtx"
	"github.com/icankeep/simplego/utils"
	"regexp"
	"strings"
)

var (
	TableNamePattern         = "(?i)create\\s+table\\s+`?([^\\s\\(`]+)`?"
	TableFieldPattern        = "(?i)\\s*`?([^`\\s]+)`?\\s+(tinyint\\(1\\)|tinyint|smallint|int|mediumint|bigint|float|double|decimal|varchar|char|tinytext|text|mediumtext|longtext|datetime|time|date|enum|set|blob|timestamp)[\\s,\\)\\(]+(.+)"
	TableFieldCommentPattern = `(?i)\s+COMMENT ['"]([^'"]+)['"]`
	TableCommentPattern      = `(?i)\s*\).+COMMENT=['"]([^'"]+)['"]`
	StructTemplate           = `%s
type %s struct {
%s}
`
	StructLineTemplate = "\t%s %s `%s` %s\n"
	StructTagTemplate  = "%s:\"%s\""
	DataTypeMap        = map[string]string{
		"tinyint(1)": "bool",
		"tinyint":    "int8",
		"smallint":   "int16",
		"mediumint":  "int32",
		"int":        "int64",
		"bigint":     "int64",
		"float":      "float32",
		"double":     "float64",
		"decimal":    "float64",
		"char":       "string",
		"varchar":    "string",
		"tinytext":   "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		"time":       "time.Time",
		"date":       "time.Time",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"enum":       "string",
		"set":        "string",
		"blob":       "string",
	}
)

type TableField struct {
	Name     string
	DataType string
	Comment  string
}

type SQLParseHandler struct {
	SQL          string
	TableName    string
	TableFields  []*TableField
	TableComment string
}

func ParseTableName(sql string) (string, bool) {
	sql = strings.TrimSpace(sql)

	reg := regexp.MustCompile(TableNamePattern)
	ret := reg.FindStringSubmatch(strings.TrimSpace(sql))
	if len(ret) < 2 {
		return "", false
	}
	return ret[1], true
}

func ParseTableFields(sql string) ([]*TableField, bool) {
	reg := regexp.MustCompile(TableFieldPattern)
	commentReg := regexp.MustCompile(TableFieldCommentPattern)
	ret := reg.FindAllStringSubmatch(sql, -1)
	if len(ret) == 0 {
		return nil, false
	}

	fields := make([]*TableField, 0)
	for _, line := range ret {
		commentMatch := commentReg.FindStringSubmatch(line[0])
		comment := utils.SafeIndexValueOrDefault(commentMatch, 1)
		fields = append(fields, &TableField{
			Name:     line[1],
			DataType: line[2],
			Comment:  comment,
		})
	}
	return fields, true
}

func ParseTableComment(sql string) (string, bool) {
	reg := regexp.MustCompile(TableCommentPattern)
	ret := reg.FindStringSubmatch(sql)
	if len(ret) == 0 {
		return "", false
	}
	return ret[1], true
}

func (h *SQLParseHandler) parse(sql string) (ok bool) {
	h.SQL = sql

	sql = strings.TrimSpace(sql)

	h.TableName, ok = ParseTableName(sql)
	if !ok {
		return false
	}

	h.TableFields, ok = ParseTableFields(sql)
	if !ok {
		return false
	}

	h.TableComment, ok = ParseTableComment(sql)
	if !ok {
		return false
	}
	return true
}

func ToGoStruct(sql string, tagTypes []string) (string, error) {

	h := &SQLParseHandler{}
	if ok := h.parse(sql); !ok {
		return "", fmt.Errorf("invalid create table SQL")
	}

	structInfo := &StructInfo{}
	// 1. 将table name转为大驼峰, Go结构体名
	structInfo.Name = UnderscoreToUpperCamelCase(h.TableName)

	// 2. 遍历字段
	fields := make([]*StructField, 0)
	for _, tableField := range h.TableFields {
		tags := make([]*StructTag, 0)
		for _, tagType := range tagTypes {
			tags = append(tags, GetTag(tagType, tableField.Name, []string{"gorm"}))
		}
		structFieldName := UnderscoreToUpperCamelCase(tableField.Name)
		fields = append(fields, &StructField{
			RealName: structFieldName,
			Name:     ProcessIDForFieldName(structFieldName),
			DataType: DataTypeMap[tableField.DataType],
			Tags:     tags,
			Comment:  tableField.Comment,
		})
	}
	structInfo.Fields = fields

	// 3. 结构体注释
	structInfo.Comment = h.TableComment

	// 4. format结构体字符串
	code := structInfo.String()
	return fmtx.FormatGoCodeOrDefault(code, code), nil
}

// ProcessIDForFieldName
// Example: XxxId => XxxID
func ProcessIDForFieldName(name string) string {
	return utils.If[string](strings.HasSuffix(name, "Id"), name[:len(name)-2]+"ID", name)
}
