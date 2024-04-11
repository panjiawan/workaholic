package putils

import "encoding/json"

func JsonMarshal(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return b
}

func JsonUnMarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}
