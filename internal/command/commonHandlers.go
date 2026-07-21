package command

import (
	"fmt"

	"github.com/saicharith2012/goredis/internal/store"
)

func handlePing(st *store.Store, args []string) string {
	if len(args) != 0 {
		response := fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 0, got %d)\r\n", CommandType(ping), len(args))
		return response
	}
	return "+PONG\r\n"
}

func handleEcho(st *store.Store, args []string) string {

	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(echo), len(args))

	}

	arg := args[0]
	return fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
}

func handleExists(st *store.Store, args []string) string {
	if len(args) == 0 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected atleast 1, got %d)\r\n", CommandType(exists), len(args))
	}

	exists := st.Exists(args)

	return fmt.Sprintf(":%d\r\n", exists)

}

func handleType(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(keyValueType), len(args))
	}

	result := st.Type(args[0])

	return fmt.Sprintf("+%s\r\n", result)
}

func handleDelete(st *store.Store, args []string) string {
	if len(args) == 0 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected atleast 1, got %d)\r\n", CommandType(delete), len(args))
	}

	deletes := st.Delete(args)

	return fmt.Sprintf(":%d\r\n", deletes)
}

func handleTTL(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(ttl), len(args))
	}

	ttl := st.TTL(args[0])

	return fmt.Sprintf(":%d\r\n", ttl)
}
