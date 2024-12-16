package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func CreateConnection() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close() // Ensure the connection is closed if the channel fails
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return conn, ch, nil
}

// ConfigureConnections sets up the queue and starts consuming messages
func ConfigureConnections(ch *amqp.Channel) {
	// Declare the queue
	q, err := ch.QueueDeclare(
		"CafeOfRestQ", // Queue name
		false,         // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name,
		"",    // Consumer name
		true,  // Auto-ack
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Listen for messages
	fmt.Println("Setting up Rabbit MQ Connection")
	// forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	// <-forever // Block forever to keep the program running
}
