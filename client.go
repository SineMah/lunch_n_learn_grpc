//go:build client
// +build client

package main

import (
	"fmt"
	"mincedmind.com/grpc/commands/average"
	"mincedmind.com/grpc/commands/hello"
	count "mincedmind.com/grpc/commands/stream"
	"os"
)

func getCommand() (string, []string) {
	commandList := os.Args[1:]

	if len(commandList) == 0 {
		fmt.Println("No command defined")
		os.Exit(1)
	}

	return commandList[0], commandList[1:]
}

func handleCommandValue(value string, args []string) {

	switch value {
	case "hello":
		hello.Do(args)
	case "count":
		count.Do(args)
	case "avg":
		average.Do()
	default:
		fmt.Printf("command %s unknown\n", value)
	}
}

func main() {
	command, args := getCommand()

	handleCommandValue(command, args)
}
