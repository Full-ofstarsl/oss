package heartbeat

import (
	"github.com/testproject/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	// with the data_server starting, connect to rabbitmq and declare(声明) a queue
	// and publish the self's listen_address to the apiservers exchange to kepp heartbeat evry 5 seconds
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
