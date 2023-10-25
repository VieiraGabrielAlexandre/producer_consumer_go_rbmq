package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Started producer ... app2")

	conn, err := amqp.Dial("amqp://producer:producer@localhost:5672/")

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

	q, err := ch.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	randomPhrase := randomdata.Paragraph()

	body := []byte(randomPhrase)

	err = ch.Publish(
		"",
		"test_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	fmt.Println("Connected to RabbitMQ")

}
