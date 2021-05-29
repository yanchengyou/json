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
	// Slice is a json SLICE
	SLICE
	// map is a json MAP
	MAP
)

func (t Type) String() string {
	switch t {
	default:
		return ""
	case Null:
		return "Null"
	case False:
		return "False"
	case Number:
		return "Number"
	case String:
		return "String"
	case True:
		return "True"
	case JSON:
		return "JSON"
	case SLICE:
		return "SLICE"
	case MAP:
		return "MAP"
	}
}

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
	// len
	Len int
	// primitive string
	primitive string
}

func Parse(jsonStr string) *Result {
	var result Result
	jsonStr = strings.TrimSpace(jsonStr)
	if strings.Count(jsonStr, "{") != strings.Count(jsonStr, "}") &&
		strings.Count(jsonStr, "[") != strings.Count(jsonStr, "]") {
		return &Result{Type: False}
	}
	if len(jsonStr) == 0 {
		return &Result{Type: Null}
	}

	result.Type = JSON
	r, err := parseToMap(jsonStr)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	jsonStr = string(js)
	result.Raw = jsonStr
	result.primitive = jsonStr
	return &result
}

func (r *Result) Get(searchStr string) *Result {
	//result, err := ParseToMap(js)
	//subList := strings.Split(searchStr, "yanchengyou")
	//if len(subList) == 1 {
	result, err := parseToMap(r.Raw)
	if err != nil {
		panic(err)
	}
	sub, ok := result[searchStr]
	if !ok {
		r.Raw = ""
		r.Type = False
		return r
	}
	switch sub.(type) {
	case []interface{}:
		r.Type = SLICE
		subLen := sub.([]interface{})
		r.Len = len(subLen)
		jsonStr, err := json.Marshal(sub)
		if err != nil {
			panic(err)
		}
		r.Raw = string(jsonStr)
	case int, float64:
		r.Type = Number
		num := sub.(float64)
		r.Num = num
		r.Raw = ""
	case string:
		r.Type = String
		str := sub.(string)
		r.Str = str
		r.Raw = ""
	case map[string]interface{}:
		r.Type = MAP
		m := sub.(map[string]interface{})
		subLen := len(m)
		jsonStr, _ := json.Marshal(sub)
		r.Raw = string(jsonStr)
		r.Len = subLen
	}
	//} else {
	//	//以后写吧
	//}
	//a, _ := mapRecMap(result, subList)
	//return a, nil
	return r
}

func (r *Result) GetOfIndex(index int) *Result {
	if !r.IsArray() {
		panic(errors.New("err for json"))
	}
	array := make([]interface{}, r.Len, r.Len)
	err := json.Unmarshal([]byte(r.Raw), &array)
	if err != nil {
		panic(err)
	}
	rStr := array[index]
	rByte, err := json.Marshal(rStr)
	if err != nil {
		panic(err)
	}
	r.Raw = string(rByte)
	r.getInterfateType(rStr)
	return r
}

func (r Result) IsArray() bool {
	if r.Type == SLICE {
		return true
	}
	return false
}

func recMap(v interface{}, search string, index int) (int, error) {
	subMap := make(map[string]interface{})
	switch v.(type) {
	case map[string]interface{}:
		subMap = v.(map[string]interface{})
		for _, vv := range subMap {
			//fmt.Printf(" ---||---%s ",ii)
			if vv != "" {
				a, _ := recMap(vv, search, index)
				if a > -1 {
					return a, nil
				}
			}

		}
	case string:
		mStr := v.(string)
		if strings.EqualFold(mStr, search) {
			return index, nil
		}
	default:
		return -2, nil
	}
	return -1, nil
}

// GetIndex 获取index
func (r *Result) GetIndex(search string) (int, error) {
	if !r.IsArray() {
		panic(errors.New("string is not a array！"))
	}
	array := make([]interface{}, r.Len, r.Len)
	err := json.Unmarshal([]byte(r.Raw), &array)
	if err != nil {
		panic(err)
	}
	for i, v := range array {
		ii, _ := recMap(v, search, i)
		if ii != -1 {
			return i, nil
		}
		//switch v.(type) {
		//case map[string]interface{}:
		//	subMap = v.(map[string]interface{})
		//}
		//for n, m := range subMap {
		//	if strings.EqualFold(n, search) {
		//		return i, nil
		//	}
		//	switch m.(type) {
		//	case string:
		//		mStr := m.(string)
		//		if strings.EqualFold(mStr, search) {
		//			return i, nil
		//		}
		//	case map[string]interface{}:
		//
		//	}
		//}
	}
	return -1, errors.New("没找到！")
}

func (r *Result) Put(key string, value interface{}, param ...interface{}) {
	if r.IsArray() {
		array := make([]interface{}, r.Len, r.Len)
		err := json.Unmarshal([]byte(r.Raw), &array)
		if err != nil {
			panic(err)
		}
		index := param[0].(int)
		str := array[index]
		arrayByte, _ := json.Marshal(array[index])
		if mapStr, ok := str.(map[string]interface{}); ok {
			mapStr[key] = value
			newStrByte, err := json.Marshal(mapStr)
			if err != nil {
				panic(err)
			}
			r.Raw = strings.ReplaceAll(r.Raw, string(arrayByte), string(newStrByte))
			r.primitive = strings.ReplaceAll(r.primitive, string(arrayByte), string(newStrByte))
			//fmt.Println(r.Raw)
			//fmt.Println(r.primitive)
		}
	} else {
		mapJson := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.Raw), &mapJson)
		if err != nil {
			panic(err)
		}
		mapJson[key] = value
		newStrByte, err := json.Marshal(mapJson)
		if err != nil {
			panic(err)
		}
		r.primitive = strings.ReplaceAll(r.primitive, r.Raw, string(newStrByte))
		r.Raw = strings.ReplaceAll(r.Raw, r.Raw, string(newStrByte))
	}
}

func (r Result) Set(value interface{}) {
	switch value.(type) {
	case string:
		strValue := value.(string)
		r.Str = strValue
	case int, float64:
		numValue := value.(float64)
		r.Num = numValue
	}
}

func (r *Result) Array() []interface{} {
	if !r.IsArray() {
		return nil
	}
	//var resultArray []Result = make([]Result, r.Len, r.Len)
	array := make([]interface{}, r.Len, r.Len)
	err := json.Unmarshal([]byte(r.Raw), &array)
	if err != nil {
		panic(err)
	}
	return array
	//for _, v := range array {
	//	switch v.(type) {
	//	case map[string]interface{}:
	//		m := sub.(map[string]interface{})
	//		subLen := len(m)
	//		jsonStr, _ := json.Marshal(sub)
	//		r.Raw = string(jsonStr)
	//		r.Len = subLen
	//	}
	//}
	//resultArray = append(resultArray, resultsub)

}

func parseToMap(js string) (map[string]interface{}, error) {
	
	var result map[string]interface{}
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

func (r *Result) getInterfateType(v interface{}) {
	switch v.(type) {
	case []interface{}:
		r.Type = SLICE
	case string:
		r.Type = String
	case int, float64:
		r.Type = Number
	case map[string]interface{}:
		r.Type = MAP
	}
}

func (r Result) String() string {

	return fmt.Sprintf("%v", r.primitive)
}
