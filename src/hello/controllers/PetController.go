package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type PetController struct {
	beego.Controller
}

func (u *PetController) ShowAllPets() {
	fmt.Println("ShowAllPets")
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/pet")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("select * from `pet`")

	var pet2 Pet
	for rows.Next() {
		err := rows.Scan(&pet2.Id, &pet2.Customer_user_id, &pet2.Name, &pet2.Species, &pet2.Breed, &pet2.Is_neuter, &pet2.Head_icon_url, &pet2.Create_time, &pet2.Update_time)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pet2)
	}

	marshal, err := json.Marshal(pet2)
	if err != nil {
		fmt.Println(err)
	}
	u.Ctx.WriteString(string(marshal))
}

func (u *PetController) GetById() {
	id, err2 := u.GetInt("id")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("getById")
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/pet")
	if err != nil {
		fmt.Println(err)
	}
	row := db.QueryRow("select * from `pet` where id=?", id)

	var pet2 Pet
	err = row.Scan(&pet2.Id, &pet2.Customer_user_id, &pet2.Name, &pet2.Species, &pet2.Breed, &pet2.Is_neuter, &pet2.Head_icon_url, &pet2.Create_time, &pet2.Update_time)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pet2)

	marshal, err := json.Marshal(pet2)
	if err != nil {
		fmt.Println(err)
	}
	u.Ctx.WriteString(string(marshal))
}

type Pet struct {
	Id               int    `json:"id"`
	Customer_user_id int    `json:"customer_User_Id"`
	Name             string `json:"name"`
	Species          int    `json:"species"`
	Breed            int    `json:"breed"`
	Is_neuter        int    `json:"is_Neuter"`
	Head_icon_url    string `json:"head_Icon_Url"`
	Create_time      string `json:"create_Time"`
	Update_time      string `json:"update_Time"`
}
