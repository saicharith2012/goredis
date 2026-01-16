package server

import (
	"strings"
)

type commandHandler func(args []string) string

var knownCommands map[string]commandHandler = map[string]commandHandler{"ping": handlePing}

func handleCommand(command string, args []string) string {
	var response string

	lowerCaseCommand := strings.ToLower(command)
	if value, ok := knownCommands[lowerCaseCommand]; ok {
		response = value(args)


	} else {
		response = "-ERR unknown command\r\n"
	}


	return response

}

func handlePing(args []string) string {
	message := "+PONG\r\n"
	return message
}
