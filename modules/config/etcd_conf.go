package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coreos/go-etcd/etcd"
)

var (
	machines = []string{"http://127.0.0.1:2379"}
)

// EtcdResponse struct data from etcd response
type EtcdResponse struct {
	Action string `json:"action"`
	Node   struct {
		CreatedIndex  float64 `json:"createdIndex"`
		Key           string  `json:"key"`
		ModifiedIndex float64 `json:"modifiedIndex"`
		Value         string  `json:"value"`
	} `json:"node"`
}

// StartEtcd beginning connection
func StartEtcd() {
	machines := []string{"http://127.0.0.1:2379"}
	client := etcd.NewClient(machines)

	if _, err := client.Set("/server_alive_w101", "alive", 0); err != nil {
		log.Fatal(err)
	}

	val, _ := client.RawGet("/server_alive_w101", true, true)

	var data EtcdResponse
	if err := json.Unmarshal(val.Body, &data); err != nil {
		fmt.Println(err)
		return
	}
}

// EtcdRawGetValue get raw or unmarshalled value from etcd cluster
func EtcdRawGetValue(key string) (string, error) {
	client := etcd.NewClient(machines)
	defer client.Close()

	val, err := client.RawGet(key, true, true)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var data EtcdResponse
	if err := json.Unmarshal(val.Body, &data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return data.Node.Value, nil

}
