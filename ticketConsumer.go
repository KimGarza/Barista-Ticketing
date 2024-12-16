package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func TicketConsumer(mqCh *amqp.Channel) {

	go func() {
		fmt.Println("Consumer started. Waiting for messages...")

		msgs, err := mqCh.Consume(
			"CafeOfRestQ", // Queue name
			"",            // Consumer name ("" lets RabbitMQ generate a unique name)
			true,          // Auto-ack (set to false if manual acknowledgment is needed)
			false,         // Exclusive
			false,         // No-local
			false,         // No-wait
			nil,           // Arguments (nil if you don't need any special options)
		)
		if err != nil {
			log.Fatalf("Failed to register consumer: %s", err)
		}

		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)

			fmt.Println("Your coffee is being made by the barista!")

			<-time.After(3 * time.Second)
			// 3 second timer complete
			fmt.Println(msgs)
		}
	}()

}
