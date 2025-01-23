package main

import (
	"github.com/testproject/interface_server/heartbeat"
	"github.com/testproject/interface_server/locate"
	"github.com/testproject/interface_server/objects"
	"github.com/testproject/interface_server/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	//心跳同步
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
