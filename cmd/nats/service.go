package nats

import (
	"github.com/fwidjaya20/demo-distributed-event-store/config"
	"github.com/nats-io/go-nats"
	"github.com/oklog/oklog/pkg/group"
	"log"
	"os"
)

func InitNatsConnection(g *group.Group) *nats.Conn {
	//cluster := config.GetEnv(config.CLUSTER_ID)
	//clientId := config.GetEnv(config.SERVICE_NAME)
	natsAddr := config.GetEnv(config.NATS_ADDR)

	//natsConn, err := stan.Connect(cluster, clientId, stan.NatsURL(natsAddr))

	natsConn, err := nats.Connect(natsAddr)

	if err != nil {
		log.Println("transport", "nats", "err", err)
		os.Exit(1)
	}

	log.Println("transport", "nats", "addr", natsAddr)

	//g.Add(
	//	func() error {
	//		log.Println("transport", "nats", "addr", natsAddr)
	//
	//		return nil
	//	},
	//	func(err error) {
	//		natsConn.Close()
	//
	//		fmt.Println(err)
	//
	//		panic(err)
	//	},
	//)

	return natsConn
}