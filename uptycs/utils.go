package uptycs

import (
	"encoding/json"
)

func SlimStructAsJsonString[T apiType](objectToSlim T, keysToDelete []string) ([]byte, error) {
	var commonKeysToDelete = []string{
		"id",
		"customerId",
		"createdAt",
		"createdBy",
		"updatedAt",
		"updatedBy",
		"links",
	}
	for _, commonKeyItem := range commonKeysToDelete {
		keysToDelete = append(keysToDelete, commonKeyItem)
	}

	rb, err := json.Marshal(objectToSlim)
	if err != nil {
		return []byte("{}"), err
	}

	var _interface interface{}
	if err := json.Unmarshal([]byte(rb), &_interface); err != nil {
		panic(err)
	}
	if m, ok := _interface.(map[string]interface{}); ok {
		for _, item := range keysToDelete {
			delete(m, item)
		}
	}
	return json.Marshal(_interface)
}
