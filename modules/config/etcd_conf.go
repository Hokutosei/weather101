package config

import (
	"log"

	"github.com/coreos/go-etcd/etcd"
)

// StartEtcd beginning connection
func StartEtcd() {
	machines := []string{"http://127.0.0.1:2379"}
	client := etcd.NewClient(machines)

	if _, err := client.Set("/server_alive_w101", "alive", 0); err != nil {
		log.Fatal(err)
	}

	val, _ := client.RawGet("/server_alive_w101", true, true)
	log.Println(string(val.Body))
}
