package jsonx

import (
	"bytes"
	"container/list"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/bytedance/sonic"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func ToJSONStr(data interface{}) (string, error) {
	result, err := sonic.Marshal(data)
	return fmt.Sprintf("%s", result), err
}

// Str2Struct Str2Struct
// 字符串转对象
func Str2Struct(source string, destination interface{}) error {
	err := sonic.Unmarshal([]byte(source), destination)
	return err
}

// list对象转数组
func List2Array(list *list.List) []interface{} {
	var len = list.Len()
	if len == 0 {
		return nil
	}
	var arr []interface{}
	for e := list.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value)
	}
	return arr
}

func ToPrettyString(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return s, e
	}
	return d, nil
}

func DecodeGbk(v interface{}, body []byte) error {
	bodyBytes, err := GbkToUtf8(body)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(bytes.NewReader(bodyBytes))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(v)
	return err
}

// Marshal marshals v into json bytes.
func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

// MarshalToString marshals v into a string.
func MarshalToString(v interface{}) (string, error) {
	data, err := Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Unmarshal unmarshals data bytes into v.
func Unmarshal(data []byte, v interface{}) error {
	if err := sonic.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func formatError(v string, err error) error {
	return fmt.Errorf("string: `%s`, error: `%w`", v, err)
}
