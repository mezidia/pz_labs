package main

import (
	"fmt"
	"lab_3/client/common"
	"lab_3/client/forums"
	"lab_3/client/users"
	"lab_3/client/dto"
)

const baseURL string = "http://localhost:8080"

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
	user := &dto.User{
		Name: "Ivan Franko",
		Mail: "ivan.franko@gmail.com",
		Password: "Djedjalyk",
		Interests: "literature",
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
