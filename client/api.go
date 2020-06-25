package client

import "fmt"

type Api struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApiList struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SpecificationVersion string `json:"specificationVersion,omitempty"`
	Apis []Api `json:"apis,omitempty"`
}

func GetAPIs(client *Client, owner string) (ApiList, error) {
	path := fmt.Sprintf("apis/%s", owner)
	result := ApiList{}
	err := client.MakeRequest("GET", path, nil, &result)

	return result, err
}
