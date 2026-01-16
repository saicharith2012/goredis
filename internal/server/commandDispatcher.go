package server

import (
	"fmt"
	"strings"
)

type commandHandler func(s *Server, args []string) string

var knownCommands map[string]commandHandler = map[string]commandHandler{"ping": handlePing, "set": handleSet, "get": handleGet}

func handleCommand(s *Server, command string, args []string) string {
	var response string

	lowerCaseCommand := strings.ToLower(command)
	if value, ok := knownCommands[lowerCaseCommand]; ok {
		response = value(s, args)
	} else {
		response = "-ERR unknown command\r\n"
	}
	return response

}

func handlePing(s *Server, args []string) string {
	message := "+PONG\r\n"
	return message
}

func handleSet(s *Server, args []string) string {
	s.store.SetValue(args[0], args[1])

	return "+OK\r\n"
}

func handleGet(s *Server, args []string) string {
	value, ok := s.store.GetValue(args[0])	

	if !ok {
		return "-1\r\n"
	}

	return fmt.Sprintf("+%s\r\n", value)
}
