package files

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

func ToStringWithError(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func ToString(filePath string) string {
	f, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(f)
}

func IsJSON(filePath string) bool {
	s, e := ToStringWithError(filePath)
	if e != nil {
		return false
	}
	return json.Unmarshal([]byte(s), new(map[string]interface{})) == nil
}

func IsXML(filePath string) bool {
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}
	return xml.Unmarshal(byteValue, new(interface{})) == nil
}
