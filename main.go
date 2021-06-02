package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	config2 "kuma.com/kumashop/config"
)

var err error
var engin *gorose.Engin

func init() {
	config, err := config2.ConfigFactory("mysql")
	if err != nil {
		panic(err)
	}
	engin, err = gorose.NewEngin(&gorose.Config{Driver: "mysql", Dsn: config.Url()})
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}

func main() {
	db := DB()
	res, err := db.Query("SELECT * FROM customer")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
