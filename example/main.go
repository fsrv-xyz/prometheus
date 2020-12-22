package main

import (
	"fmt"

	Prometheus "./prometheus"
)

func main() {
	client := Prometheus.Client{Endpoint: "http://10.1.50.9:9090"}
	data := client.Query("up").Result()
	fmt.Println(data)
}
