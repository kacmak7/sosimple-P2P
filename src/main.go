package main 

import (
	//"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	
)

func launchServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test).Methods("GET")
	//router.HandleFunc("/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	//port := flag.String("port", "")

	go launchServer()

	initializeNode()
	send("HI HELLOooo")

	time.Sleep(1000 * time.Second)
}
