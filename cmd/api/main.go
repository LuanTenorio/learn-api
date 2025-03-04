package main

import (
	_ "github.com/LuanTenorio/learn-api/docs"
	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/server"
)

// @title			Learn
// @version		1.0
// @description	API of the Learn project, which aims to organize and analyze studies
// @BasePath		/api
func main() {
	conf := config.GetConfig()
	var db database.Database = database.NewPGDatabase(&conf)
	server.NewEchoServer(&conf, db).Start()
}
