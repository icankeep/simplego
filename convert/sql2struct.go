package main

import (
	"fmt"
	"regexp"
	"strings"
)

const sql = `
CREATE TABLE Persons (
    PersonID int COMMENT '123',
    LastName varchar(255),
    FirstName varchar(255),
    Address varchar(255),
    City varchar(255)
) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
`

type Person struct {
	a string
	b int
}

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

const (
	TableNamePattern    = "(?i)create\\s+table\\s+`?([^\\s\\(]+)`?"
	TableFieldPattern   = "(?i)\\s*`?([^`\\s]+)`?\\s+(tinyint|smallint|int|mediumint|bigint|float|double|decimal|varchar|char|tinytext|text|mediumtext|longtext|datetime|time|date|enum|set|blob|timestamp).*(comment\\s+['\"]([^'\"]+)['\"])"
	TableCommentPattern = `(?i)\s*\).+COMMENT=['"]([^'"]+)['"]`
)

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
	ret := reg.FindAllStringSubmatch(sql, -1)
	if len(ret) == 0 {
		return nil, false
	}

	fields := make([]*TableField, 0)
	for _, line := range ret {
		fields = append(fields, &TableField{
			Name:     line[1],
			DataType: line[2],
			Comment:  line[3],
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

func (h *SQLParseHandler) ToGoStruct(sql string) error {
	if ok := h.parse(sql); !ok {
		return fmt.Errorf("invalid create table SQL")
	}

	// 1. 将table name转为大驼峰, Go结构体名

	// 2. 遍历字段

	// 3. format结构体字符串
}

func main() {
	s := strings.TrimSpace(sql)
	ParseTableComment(s)
}

const StructTemplate = `
type {tableName} struct {
	
}
`

//
//func SQL2Struct(sql string) string {
//
//}
