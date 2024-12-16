package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

func Order(writer *bufio.Writer, reader *bufio.Reader, mqCh *amqp.Channel) {
	coffee, _ := CoffeeItem(writer, reader)
	temp := "hot"
	isHot, _ := HotIced(reader, writer)
	if !isHot {
		temp = "iced"
	}
	milk, _ := MilkType(reader, writer)

	// rabbit mq submit order ticket
	// create json object of ticket
	// stringify it
	// Publish a message
	body := fmt.Sprintf("Your order is ready!\nA(n) %s %s made with %s", temp, coffee, milk)
	err := mqCh.Publish(
		"",            // Exchange (default exchange)
		"CafeOfRestQ", // Routing key (queue name)
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf("[x] Sent\n")

}

func CoffeeItem(writer *bufio.Writer, reader *bufio.Reader) (string, error) {
	itemNum := 0
	menuItems := []string{
		"Undying Espresso",
		"Necromancer's Macchiato",
		"Tolman's Hope",
		"Piety's Spiced Latte",
		"Exile's Energizer",
		"Solaris Dark Mocha",
		"Lunaris White Mocha",
		"Riverways Cold Brew",
	}

	validInput := false
	for !validInput {
		for item := range menuItems {
			itemNum++
			writer.WriteString(fmt.Sprintf("%d) %s\n", itemNum, menuItems[item]))
		}
		writer.Flush()

		input, err := GetInput(reader)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to get input: %w", err))
			return "", err
		}

		inputNum, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert input to num, %w", err))
			return "", err
		}

		if inputNum >= 1 && inputNum <= itemNum {
			writer.WriteString(fmt.Sprintf("%s it is!\n", menuItems[inputNum-1]))
			writer.Flush()
			return menuItems[inputNum-1], nil
		} else {
			itemNum = 0
			fmt.Println("Sorry, what?")
		}
	}
	return "", nil
}

func HotIced(reader *bufio.Reader, writer *bufio.Writer) (bool, error) {

	validInput := false
	for !validInput {
		writer.WriteString("1) Hot\n")
		writer.WriteString("2) Iced\n")
		writer.Flush()
		input, err := GetInput(reader)
		if err != nil {
			return true, err
		}

		if strings.TrimSpace(input) == "1" {
			return true, nil
		} else if strings.TrimSpace(input) == "2" {
			return false, nil
		} else {
			writer.WriteString("Sorry, what?\n")
			writer.Flush()
		}
	}

	return false, nil
}

func MilkType(reader *bufio.Reader, writer *bufio.Writer) (string, error) {
	itemNum := 0
	milkTypes := []string{
		"2%",
		"Whole Milk",
		"Half and half",
		"Whole Milk",
		"Roah Milk",
		"Hazelnut Milk",
	}

	validInput := false
	for !validInput {
		for item := range milkTypes {
			itemNum++
			writer.WriteString(fmt.Sprintf("%d) %s\n", itemNum, milkTypes[item]))
		}
		writer.Flush()

		input, err := GetInput(reader)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to get input: %w", err))
			return "", err
		}

		inputNum, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert input to num, %w", err))
			return "", err
		}

		if inputNum >= 1 && inputNum <= itemNum {
			writer.WriteString(fmt.Sprintf("%s it is!\n", milkTypes[inputNum-1]))
			writer.Flush()
			return milkTypes[inputNum-1], nil
		} else {
			itemNum = 0
			fmt.Println("Sorry, what?")
		}
	}
	return "", nil
}
