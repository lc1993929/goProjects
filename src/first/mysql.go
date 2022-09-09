package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	open, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/pet")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(open)
	row := open.QueryRow("select * from `order`")
	var id, customer_user_id, store_id, is_pay, status int
	var price float64
	var create_time, update_time string
	err = row.Scan(&id, &customer_user_id, &store_id, &price, &is_pay, &status, &create_time, &update_time)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, customer_user_id, store_id, price, is_pay, status, create_time, update_time)
}
