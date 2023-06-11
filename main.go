package main

import (
	"collection/httpserv"
	"collection/infrastructure"
	"collection/logs"
)

func init() {
	infrastructure.InitConfig()
	logs.InitLogger()
}

func main() {
	infrastructure.InitTimeZone()
	infrastructure.InitDB()
	httpserv.Run()
}
