package api

import (
	"corona/api/controllers"
)

var server = controllers.Server{}

func Run() {

	server.Initialize()
	server.Run(":8000")

}
