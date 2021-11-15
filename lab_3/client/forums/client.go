package forums

import (
	"lab_3/client/common"
)

//return method get forums using http methods from package common

func GetForums(client *common.Client) (string, error) {
	return client.Get("/forums")
}
