package uptycs

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

var validate = validator.New()

func SlimStructAsJSONString[T iAPIType](objectToSlim T, keysToDelete []string) ([]byte, error) {
	var commonKeysToDelete = []string{
		"id",
		"customerId",
		"createdAt",
		"createdBy",
		"updatedAt",
		"updatedBy",
		"links",
		"lock",   //"lock"  of event rule should be a boolean" api response if included
		"custom", // You could set this on custom rules to true and cause issues requiring support
	}
	keysToDelete = append(keysToDelete, commonKeysToDelete...)

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

func doDelete[T iAPIType](c *Client, apiObject T, endpointStr string) (T, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.HostURL, endpointStr, apiObject.GetID()), nil)
	if err != nil {
		return apiObject, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return apiObject, err
	}

	return apiObject, nil
}

func doGetMany[T iAPITypes](c *Client, apiObject T, endpointStr string) (T, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, endpointStr), nil)
	if err != nil {
		return apiObject, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return apiObject, err
	}

	foundItems := make([]T, 1)

	err = json.Unmarshal(body, &foundItems[0])
	if err != nil {
		return apiObject, err
	}

	return foundItems[0], nil
}

func doGet[T iAPIType](c *Client, apiObject T, endpointStr string) (T, error) {
	urlStr := fmt.Sprintf("%s/%s/%s", c.HostURL, endpointStr, apiObject.GetID())
	if len(apiObject.GetID()) == 0 {
		urlStr = fmt.Sprintf("%s/%s", c.HostURL, endpointStr)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return apiObject, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return apiObject, err
	}

	foundItems := make([]T, 1)

	if len(apiObject.GetID()) == 0 {
		// Attempt to get by name using doGetMany() and GetName() on each
		panic("TODO: support get by name if len(GetID()) is 0")
	} else {
		err = json.Unmarshal(body, &foundItems[0])
		if err != nil {
			return apiObject, err
		}
	}

	return foundItems[0], nil
}

func doCreate[T iAPIType](c *Client, apiObject T, endpointStr string, additionalKeysToDelete []string) (T, error) {
	err := validate.Struct(apiObject)
	if err != nil {
		return apiObject, err
	}

	slimmedObj, err := SlimStructAsJSONString(apiObject, append(apiObject.KeysToDelete(), additionalKeysToDelete...))
	if err != nil {
		return apiObject, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.HostURL, endpointStr),
		strings.NewReader(string(slimmedObj)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return apiObject, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return apiObject, err
	}

	newAPIObject := make([]T, 1)
	err = json.Unmarshal(body, &newAPIObject[0])
	if err != nil {
		return apiObject, err
	}
	return newAPIObject[0], nil
}

func doUpdate[T iAPIType](c *Client, apiObject T, endpointStr string, additionalKeysToDelete []string) (T, error) {
	err := validate.Struct(apiObject)
	if err != nil {
		return apiObject, err
	}
	slimmedObj, err := SlimStructAsJSONString(apiObject, append(apiObject.KeysToDelete(), additionalKeysToDelete...))
	if err != nil {
		return apiObject, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/%s/%s", c.HostURL, endpointStr, apiObject.GetID()),
		strings.NewReader(string(slimmedObj)),
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return apiObject, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return apiObject, err
	}

	newAPIObject := make([]T, 1)
	err = json.Unmarshal(body, &newAPIObject[0])
	if err != nil {
		return apiObject, err
	}

	return newAPIObject[0], nil
}
