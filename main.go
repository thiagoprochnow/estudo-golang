package main

import (
	"fmt"

	"estudo.com/helpers"
	"github.com/cdt-baas/gocore/httpserver/irisapi"
)

func main() {
	var serverPort = helpers.GetDefaultPort()
	fmt.Println("Server port: ", serverPort)

	_app := irisapi.IrisApi()
	_app.RunServer(serverPort)
}
