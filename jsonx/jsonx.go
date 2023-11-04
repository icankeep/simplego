package jsonx

import (
	"bytes"
	"encoding/json"
	"strings"
)

func MustToJSON(v interface{}) string {
	if value, err := ToJSON(v); err != nil {
		panic(err)
	} else {
		return value
	}
}

func ToFormatJSON(v interface{}, indent int) (string, error) {
	bs, err := json.MarshalIndent(v, "", strings.Repeat(" ", indent))
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func ToJSON(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func MustJSONToAny(s string, a interface{}) error {
	if err := JSONToAny(s, a); err != nil {
		panic(err)
	}
	return nil
}

func JSONToAny(s string, a interface{}) error {
	d := json.NewDecoder(bytes.NewReader([]byte(s)))
	d.UseNumber()

	return d.Decode(a)
}

func FormatJSON(s string, indent int) (string, error) {
	var a interface{}
	err := JSONToAny(s, &a)
	if err != nil {
		return "", err
	}

	formatJSON, err := ToFormatJSON(a, indent)
	if err != nil {
		return "", err
	}
	return formatJSON, nil
}
