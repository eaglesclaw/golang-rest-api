package main

import (
	"Youtube-Projects/golang-rest-api/httpserver/pkg/ui"

	"github.com/eaglesclaw/golang-rest-api/pkg/api"
	"github.com/eaglesclaw/golang-rest-api/pkg/database"
	"github.com/eaglesclaw/golang-rest-api/pkg/ui"

	"fmt"
)



func main(){
	fmt.Println("......main")


	api.Api()
	database.Database()
	ui.Ui()
}