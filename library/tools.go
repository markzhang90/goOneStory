package library

import "encoding/json"

func Json2Map(input string)(map[string]interface{}, error)  {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(input), &result); err != nil {
		return nil, err
	}
	return result, nil
}