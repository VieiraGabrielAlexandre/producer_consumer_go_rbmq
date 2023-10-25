package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	fmt.Println("Started consumer ... app1")
	conn, err := amqp.Dial("amqp://consumer:consumer@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer ch.Close()

	forever := make(chan bool)

	fileName := "received.txt"

	go func() {
		msgs, _ := ch.Consume(
			"test_queue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		for d := range msgs {
			fmt.Println("Received message: ", string(d.Body))

			file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

			if err != nil {
				fmt.Println(err)
				panic(1)
			}

			defer file.Close()

			_, err = file.WriteString(string(d.Body) + "\n")

			if err != nil {
				fmt.Println(err)
				panic(1)
			}
		}
	}()

	fmt.Println("Waiting for messages ...")
	<-forever

	fmt.Println("Connected to RabbitMQ")
}
