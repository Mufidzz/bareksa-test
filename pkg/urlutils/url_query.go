package urlutils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func EncodeStruct(value interface{}) (encodedString string, err error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	encodedString = base64.StdEncoding.EncodeToString(jsonValue)
	encodedString = strings.ReplaceAll(encodedString, "+", ".")
	encodedString = strings.ReplaceAll(encodedString, "/", "_")
	encodedString = strings.ReplaceAll(encodedString, "=", "-")

	return encodedString, nil
}

func DecodeEncodedString(encodedFilterString string, v interface{}) (err error) {
	encodedFilterString = strings.ReplaceAll(encodedFilterString, ".", "+")
	encodedFilterString = strings.ReplaceAll(encodedFilterString, "_", "/")
	encodedFilterString = strings.ReplaceAll(encodedFilterString, "-", "=")

	decodedByte, err := base64.StdEncoding.DecodeString(encodedFilterString)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decodedByte, &v)
	if err != nil {
		return err
	}

	return nil
}
