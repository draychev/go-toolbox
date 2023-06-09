package json

import (
	"encoding/json"
	"fmt"
)

// PrettyJSON will Unmarshal and Marshall again with Indent, so it is human-readable
func PrettyJSON(js []byte, prefix string) ([]byte, error) {
	var jsonObj interface{}
	err := json.Unmarshal(js, &jsonObj)
	if err != nil {
		return nil, fmt.Errorf("could not Unmarshal a byte array: %w", err)
	}
	return json.MarshalIndent(jsonObj, prefix, "    ")
}
