package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Print("pokedex >")

		var input string
		cmd, err := fmt.Scanln(&input)
		
		cmdStr := fmt.Sprintf("Command: %v", cmd)
		inputStr := fmt.Sprintf("Input: %s", input)

		fmt.Println(cmdStr)
		fmt.Println(inputStr)

		if err == nil {
			switch input {
				case "help":
					fmt.Println("You selected Help")
				case "exit":
					fmt.Println("You selected Exit")
					return
				default:
					fmt.Println("Command not recognised")
			}
		} else {
			print(err)
		}
	}
}