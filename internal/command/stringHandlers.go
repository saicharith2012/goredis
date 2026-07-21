package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/saicharith2012/goredis/internal/store"
)

func handleSet(st *store.Store, args []string) string {
	if len(args) == 2 {
		st.Set(args[0], args[1], 0)
	} else if len(args) == 4 && strings.ToUpper(args[2]) == "EX" {
		ttl, err := strconv.Atoi(args[3])

		if err != nil {
			return fmt.Sprintf("-ERR %s", err.Error())
		}

		if ttl <= 0 {
			return fmt.Sprintf("-ERR ttl must be greater than 0 (expected ttl > 0, got %d)\r\n", ttl)
		}

		st.Set(args[0], args[1], int64(ttl))
	} else {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 2 or 4 with ttl, got %d)\r\n", CommandType(set), len(args))
	}

	return "+OK\r\n"
}

func handleGet(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(get), len(args))
	}

	value, found, err := st.Get(args[0])

	if !found {
		return "$-1\r\n"
	}

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}

func handleIncrement(st *store.Store, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command (expected 1, got %d)\r\n", CommandType(incr), len(args))
	}

	val, err := st.Incr(args[0])

	fmt.Println(val, err)

	if errors.Is(err, store.ErrWrongType) {
		return "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	}

	if errors.Is(err, store.ErrInvalidInteger) {
		return "-ERR value is not an integer or out of range\r\n"
	}

	return fmt.Sprintf(":%d\r\n", val)
}