package main

import (
	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/server"
)

func main() {
	conf := config.GetConfig()
	var db database.Database = database.NewPostgresDatabase(&conf)
	server.NewEchoServer(&conf, db).Start()
}
