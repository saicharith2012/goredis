package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/saicharith2012/goredis/internal/store"
)


func  handleLPush(st *store.Store, args []string) string {
	if len(args) < 2 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected atleast 2, got %d)\r\n", CommandType(lpush), len(args))
	}

	count, err := st.LPush(args[0], args[1:])

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf(":%d\r\n", count)
}

func handleRPush(st *store.Store, args []string) string {
	if len(args) < 2 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected atleast 2, got %d)\r\n", CommandType(rpush), len(args))
	}

	count, err := st.RPush(args[0], args[1:])

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf(":%d\r\n", count)
}

func handleLLen(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(llen), len(args))
	}

	length, err := st.LLen(args[0])

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf(":%d\r\n", length)
}

func handleLRange(st *store.Store, args []string) string {
	if len(args) != 3 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 3, got %d)\r\n", CommandType(lrange), len(args))
	}

	key := args[0]
	start, err := strconv.Atoi(args[1])

	if err != nil {
		return fmt.Sprintf("-ERR invalid start boundary (expected int, got %s)\r\n", args[1])
	}

	end, err := strconv.Atoi(args[2])

	if err != nil {
		return fmt.Sprintf("-ERR invalid end boundary (expected int, got %s)\r\n", args[1])
	}

	list, err := st.LRange(key, start, end)

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	response := []string{}

	response = append(response, fmt.Sprintf("*%d\r\n", len(list)))

	for _, listItem := range list {
		response = append(response, fmt.Sprintf("$%d\r\n%s\r\n", len(listItem), listItem))
	}

	return strings.Join(response, "")
}

func handleLPop(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(lpop), len(args))
	}

	value, err := st.LPop(args[0])

	if errors.Is(err, store.ErrNotFound) {
		return "$-1\r\n"
	}

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}

func handleRPop(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(rpop), len(args))
	}

	value, err := st.RPop(args[0])

	if errors.Is(err, store.ErrNotFound) {
		return "$-1\r\n"
	}

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}
