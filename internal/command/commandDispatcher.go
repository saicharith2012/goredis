package command

import (
	"fmt"
	"strings"

	"github.com/saicharith2012/goredis/internal/store"
)

type CommandHandler func(st *store.Store, args []string) string

type CommandType string

const (
	ping         CommandType = "PING"
	echo         CommandType = "ECHO"
	set          CommandType = "SET"
	get          CommandType = "GET"
	keyValueType CommandType = "TYPE"
	exists       CommandType = "EXISTS"
	delete       CommandType = "DEL"
	ttl          CommandType = "TTL"
	incr         CommandType = "INCR"
	lpush        CommandType = "LPUSH"
	rpush        CommandType = "RPUSH"
	llen         CommandType = "LLEN"
	lrange       CommandType = "LRANGE"
	lpop         CommandType = "LPOP"
	rpop         CommandType = "RPOP"
)

var commandHandlers = map[CommandType]CommandHandler{
	ping:         handlePing,
	echo:         handleEcho,
	set:          handleSet,
	get:          handleGet,
	keyValueType: handleType,
	exists:       handleExists,
	delete:       handleDelete,
	ttl:          handleTTL,
	incr:         handleIncrement,
	lpush:        handleLPush,
	rpush:        handleRPush,
	llen:         handleLLen,
	lrange:       handleLRange,
	lpop:         handleLPop,
	rpop:         handleRPop,
}

func HandleCommand(tokens []string, st *store.Store) string {
	command := tokens[0]
	args := tokens[1:]

	upperCaseCommand := CommandType(strings.ToUpper(command))

	if _, known := commandHandlers[upperCaseCommand]; known {
		return commandHandlers[upperCaseCommand](st, args)
	} else {
		return fmt.Sprintf("-ERR unknown command '%s'\r\n", command)
	}
}
