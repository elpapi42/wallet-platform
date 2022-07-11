package eventstore

import (
	"fmt"
	"log"
	"sync"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

var client *esdb.Client

var lock = &sync.Mutex{}

func GetClient() *esdb.Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			settings, err := esdb.ParseConnectionString(
				"esdb://localhost:2111,localhost:2112,localhost:2113?tls=false&keepAliveTimeout=10000&keepAliveInterval=10000",
			)

			if err != nil {
				panic(err)
			}

			client, err = esdb.NewClient(settings)

			if err != nil {
				panic(err)
			}

			log.Println("eventstore client created")
		} else {
			fmt.Println("eventstore client already created")
		}
	}

	return client
}
