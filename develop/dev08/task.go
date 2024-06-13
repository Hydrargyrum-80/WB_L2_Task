package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	strScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !strScanner.Scan() {
			break
		}
		input := strScanner.Text()
		commands := strings.Split(input, " | ")
		for _, command := range commands {
			args := strings.Split(command, " ")
			switch args[0] {
			case "cs":
			case "ls":
			case "pwd":
			case "echo":
			case "kill":
			case "ps":
			case "quit":
				os.Exit(0)
			default:
				fmt.Println("Command not found")
				return
			}
			cmd := exec.Command(args[0], args[1:]...)
			err := cmd.Err
			if err != nil {
				log.Fatal(err.Error())
			}
			bytes, err := cmd.Output()
			if err != nil {
				log.Fatal(err.Error())
			}
			str := string(bytes)
			fmt.Println(str[:len(str)-1])
		}
	}
}
