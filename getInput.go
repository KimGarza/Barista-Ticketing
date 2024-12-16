package main

import (
	"bufio"
	"fmt"
)

func GetInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("error reading input: %w", err))
		return "", err
	}
	return input, nil
}
