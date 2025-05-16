package main

import (
	"fmt"
	"listener-service/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	// Try to connect to rabbit mq.
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")

	// Start listening for messages

	log.Println("Listening for and consuming RabbitMQ messages...")

	// create a consumer

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue, and consume the events.

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)

	}

	//
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// do not conitune until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println("Error while connecting to RabbitMQ: ", err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
	}
	return connection, nil
}
