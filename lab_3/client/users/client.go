package users

import (
	"github.com/mezidia/pz_labs/lab_3/client/common"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
)

//return method registrate user using http method from common
func RegistrateUser(client *common.Client, userInfo *dto.User) (*dto.RegistrateUserResponse, error) {
	return client.Post("/users", userInfo)
}
