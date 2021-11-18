package main

import (
	"fmt"
	"github.com/mezidia/pz_labs/lab_3/client/common"
	"github.com/mezidia/pz_labs/lab_3/client/forums"
	"github.com/mezidia/pz_labs/lab_3/client/users"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
)

const baseURL string = "http://localhost:8081"

func main() {
	client := &common.Client{
		BaseURL: baseURL,
	}

	//scenario get all forums
	fmt.Println("=== Scenario 1 ===")
	fmt.Println("All forums: ")
	forums, err := forums.GetForums(client)
	if err != nil {
		fmt.Println(err)
		return
 	}
	fmt.Println(forums)

	//scenario registrate user
	interests := make([]string, 3)
	interests[0] = "literature"
  interests[1] = "politics"
	interests[2] = "cars"
	user := &dto.User{
		Name: "Ivan Franko",
		Mail: "ivan.franko@gmail.com",
		Password: "Djedjalyk",
		Interests: interests,
	}
	fmt.Println("=== Scenario 2 ===")
	fmt.Println("Registrate user: ")
	success, err := users.RegistrateUser(client, user)
	if err != nil {
		fmt.Println(err)
		return
 	}
	fmt.Println(success)
	
}
