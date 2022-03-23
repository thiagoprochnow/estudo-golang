package main

import (
	"fmt"
	"net/http"
	"strconv"

	"estudo.com/controllers"
	"estudo.com/helpers"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {
	var serverPort = helpers.GetDefaultPort()
	fmt.Println("Server port: ", serverPort)

	http.HandleFunc("/kafka", controllers.Get_kafka)

	http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)
}
