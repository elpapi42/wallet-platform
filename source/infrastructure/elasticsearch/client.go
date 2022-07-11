package elasticsearch

import (
	"fmt"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

var lock = &sync.Mutex{}

func GetClient() *elasticsearch.Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			var err error
			client, err = elasticsearch.NewClient(elasticsearch.Config{
				Addresses: []string{
					"http://localhost:9200",
					"http://localhost:9201",
					"http://localhost:9202",
				},
			})

			if err != nil {
				panic(err)
			}

			log.Println("elasticsearch client created")
		} else {
			fmt.Println("elasticsearch client already created")
		}
	}

	return client
}
