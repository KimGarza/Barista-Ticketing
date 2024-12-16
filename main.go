package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	mqConn, mqCh, err := CreateConnection()
	if err != nil {
		log.Fatalf("Error setting up RabbitMQ: %s", err)
	}
	defer mqConn.Close() // Close the connection when done
	defer mqCh.Close()   // Close the channel when done

	// Configure consumer
	ConfigureConnections(mqCh)
	TicketConsumer(mqCh)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	validInput := false

	for !validInput {
		writer.WriteString("1)	Order\n")
		writer.WriteString("2)	History\n")
		writer.Flush()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(fmt.Errorf("failed to get input: %w", err))
		}

		input = strings.TrimSpace(input)
		if input == "1" {
			Order(writer, reader, mqCh)
			validInput = true
		} else if input == "2" {
			Lore(writer)
		} else {
			writer.WriteString("Sorry, what?\n")
		}
	}
}
