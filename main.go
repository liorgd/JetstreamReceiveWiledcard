package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

const (
	ordersSubject = "ORDERS.*"
)

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

func main() {
	log.Println("Connect to NATS")
	nc, _ := nats.Connect("demo.nats.io")
	log.Println("Creates JetStreamContext")
	js, err := nc.JetStream()
	checkErr(err)
	log.Printf("Create durable consumer monitor on subject:%q", ordersSubject)
	js.Subscribe(ordersSubject, func(msg *nats.Msg) {
		msg.Ack()
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		checkErr(err)
		log.Printf("Subscriber fetched msg.Data:%s from subSubjectName:%q", string(msg.Data), msg.Subject)
	}, nats.Durable("monitor"), nats.ManualAck())

	runtime.Goexit()

}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
