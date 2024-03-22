package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var data map[string]interface{}

	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		fmt.Println("Error: invalid json:", err)
		return
	}

	for key, value := range data {
		if key == "_public_key" {
			continue
		}
		fmt.Printf("%s=%v\n", key, value)
	}
	fmt.Println()
}
