package main

import (
	. "./controller"
	. "./config"
	. "./mapper"
)

var config = Config{}
var dao = BookDAO{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	Route()
}
