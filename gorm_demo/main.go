package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study.go/gorm_demo/config"
)

func main() {
	var mysqlConfig = config.Config.Mysql
	var ptr = &mysqlConfig.Username
	*ptr = "hahaha"
	fmt.Printf("Mysql addr --> %p\n", &(config.Config.Mysql.Username))
	fmt.Printf("mysqlConfig addr --> %p\n",&(mysqlConfig.Username))
	fmt.Println(config.Config.Mysql)
	dsn := "root:abc123@tcp(192.168.131.125:3306)/iplaymtg?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user user
	db.Raw("select id, username, head, fire from user where id = ?", 1453295).Scan(&user)
	fmt.Println(user)

}

type user struct {
	Id       int
	Username string
	Head     string
	Fire     uint
}
