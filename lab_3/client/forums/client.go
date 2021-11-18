package forums

import (
	"github.com/mezidia/pz_labs/lab_3/client/common"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
)

//return method get forums using http methods from package common
func GetForums(client *common.Client) ([]dto.Forum, error) {
	return client.Get("/forums")
}
