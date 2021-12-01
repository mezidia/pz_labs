package main

import (
	"fmt"

	"github.com/mezidia/pz_labs/lab_3/client/common"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
	"github.com/mezidia/pz_labs/lab_3/client/forums"
	"github.com/mezidia/pz_labs/lab_3/client/users"
)

const baseURL string = "http://localhost:8081"

func main() {
	client := &common.Client{
		BaseURL: baseURL,
	}

	//scenario get all forums
	fmt.Println("=== Scenario 1 ===")
	fmt.Println("All forums: ")
	forums1, err := forums.GetForums(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(forums1)

	//scenario registrate user
	interests := make([]string, 3)
<<<<<<< HEAD
	interests[0] = "literature"
	interests[1] = "politics"
	interests[2] = "War"
	user := &dto.User{
		UserName:  "Ivan Franko",
		UserType:  0,
		UserMail:  "ivan.franko@gmail.com",
		Password:  "Djedjalyk",
=======
	interests[0] = "Yagnyata"
	interests[1] = "flowers1"
	interests[2] = "War"
	user := &dto.User{
		UserName:      "Taras Shevchenko",
		UserType: 0,
		UserMail:      "taras.shevchenko@gmail.com",
		Password:  "Slava Ukrayini",
>>>>>>> 71fdbd32bcccaad00ce7bdd3e0318923ee770cae
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

	//scenario get all forums
	fmt.Println("=== Scenario 1 ===")
	fmt.Println("All forums: ")
	forums2, err := forums.GetForums(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(forums2)

}
