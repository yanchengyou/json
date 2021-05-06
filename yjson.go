package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Type int

const (
	// Null is a null json value
	Null Type = iota
	// False is a json false boolean
	False
	// Number is json number
	Number
	// String is a json string
	String
	// True is a json true boolean
	True
	// JSON is a raw block of JSON
	JSON
)

type Result struct {
	// Type is the json type
	Type Type
	// Raw is the raw json
	Raw string
	// Str is the json string
	Str string
	// Num is the json number
	Num float64
	// Index of raw value in original json, zero means index unknown
	Index int
}

func Parse(json string) *Result {
	var result Result
	json = strings.TrimSpace(json)
	if strings.Count(json, "{") != strings.Count(json, "}") &&
		strings.Count(json, "[") != strings.Count(json, "]") {
		return &Result{Type: False}
	}
	if len(json) == 0 {
		return &Result{Type: Null}
	}

	result.Type = JSON
	result.Raw = json
	return &result
}

func (r *Result) Get(searchStr string) *Result {
	//result, err := ParseToMap(js)
	subList := strings.Split(searchStr, ".")
	if len(subList) == 1 {
		result, err := parseToMap(r.Raw)
		if err != nil {
			panic(err)
		}
		sub := result[subList[0]]
		jsonStr, err := json.Marshal(sub)
		if err != nil {
			panic(err)
		}
		r.Raw = string(jsonStr)
		switch sub.(type) {
		case []interface{}:
			fmt.Println("[]interface{}")
		case string:
			r.Type = String
		}
	}
	//a, _ := mapRecMap(result, subList)
	//return a, nil
	return r
}

var result map[string]interface{}

func parseToMap(js string) (map[string]interface{}, error) {

	err := json.Unmarshal([]byte(js), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Get(js, searchStr string) (interface{}, error) {

	result, err := parseToMap(js)
	if err != nil {
		return nil, err
	}
	subList := strings.Split(searchStr, ".")
	if len(subList) == 1 {
		return result[subList[0]], nil
	}
	a, _ := mapRecMap(result, subList)
	return a, nil
}

func mapRecMap(s map[string]interface{}, sl []string) (interface{}, error) {
	for i, v := range sl {
		if len(sl) <= 1 {
			return s[v], nil
		} else {
			if sok, ok := s[v].(map[string]interface{}); ok {
				slNow := sl[i+1:]
				return mapRecMap(sok, slNow)
			}
		}
	}
	return "err", nil
}

func PutStr(js string, searchStr string, subKey string, searchValue string, key string,
	value interface{}) (map[string]interface{}, error) {

	searchResult, _ := Get(js, searchStr)
	switch searchResult.(type) {
	case []interface{}:
		if m, ok := searchResult.([]interface{}); ok {
			for _, v := range m {
				if sub, ok := v.(map[string]interface{}); ok {
					subValue := sub[subKey].(string)
					if strings.EqualFold(subValue, searchValue) {
						sub[key] = value
					}
				}
			}
		}
		return result, nil
	default:
		fmt.Println(result)
		return nil, errors.New("断言error")
	}
}
