package users

import (
	"github.com/mezidia/pz_labs/lab_3/client/common"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
)

//return method registrate user using http method from common
<<<<<<< HEAD
func RegistrateUser(client *common.Client, userInfo *dto.User) (*dto.RegistrateUserResponse, error) {
	return client.Post("/users/registrateUser", userInfo)
=======
func RegistrateUser(client *common.Client, userInfo *dto.User) (int, error) {
	return client.Post("/users", userInfo)
>>>>>>> 71fdbd32bcccaad00ce7bdd3e0318923ee770cae
}
