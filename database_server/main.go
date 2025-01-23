package main

import (
	"github.com/testproject/database_server/heartbeat"
	"github.com/testproject/database_server/locate"
	"github.com/testproject/database_server/objects"
	temp "github.com/testproject/database_server/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
