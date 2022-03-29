package main

import (
	"fmt"
	"net/http"
	"strconv"

	"estudo.com/controllers"
	"estudo.com/helpers"
	"estudo.com/services"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {
	var serverPort = helpers.GetDefaultPort()
	fmt.Println("Server port: ", serverPort)
	go services.GetKafkaContent()

	http.HandleFunc("/transaction", controllers.GetTransaction)

	http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)
}
