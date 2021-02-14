package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var str = "This is a variable of an other .go file in main package"

func dance() {
	fmt.Println("hello，let‘s dancing")
}

func dataCon() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/steven")
	if err != nil {
		log.Fatal(err)
	}

	// 查询
	rows, err := db.Query("select id,user,sex,pwd,createtime,mtime from t_user order by id asc")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var user string
		var pwd string
		var sex int
		var createtime string
		var mtime string
		err = rows.Scan(&id, &user, &pwd, &sex, &createtime, &mtime)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, user, pwd, sex, createtime, mtime)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
