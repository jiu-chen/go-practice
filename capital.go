package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roam"
	countryCapitalMap["Japan"] = "Tokoy"
	countryCapitalMap["India"] = "Newdeli"

	for country := range countryCapitalMap {
		fmt.Println(country, "capital is ", countryCapitalMap[country])
	}

	// 查看元素是否在集合中
	if capital, ok := countryCapitalMap["America"]; !ok {
		fmt.Println("America is not in map.")
	} else {
		fmt.Println("America capital is ", capital)
	}

	//  插入数据库
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range countryCapitalMap {
		result, err := db.Exec("INSERT INTO countryCapital(country, capital) VALUES(?,?)", k, v)
		if err != nil {
			fmt.Println(result, err)
		}
	}

}

/*
创建数据库: CREATE DATABASE test;

创建表：
CREATE TABLE `countryCapital` (
	id int primary KEY AUTO_INCREMENT,
	country varchar(32),
	capital varchar(32)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

// 下载驱动包:
// go get github.com/go-sql-driver/mysql
