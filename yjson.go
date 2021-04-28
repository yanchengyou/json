package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var result map[string]interface{}

func ParseToMap(js string) (map[string]interface{}, error) {

	err := json.Unmarshal([]byte(js), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Get(js, searchStr string) (interface{}, error) {

	result, err := ParseToMap(js)
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

func PutStr(js string, searchStr string, subKey string, searchValue string, key string, value interface{}) (map[string]interface{}, error) {

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
