package server

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func respParser(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, "*") {
		return nil, fmt.Errorf("expected array, got %s", line)
	}

	count, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, count)

	for i := 0; i < count; i++ {
		line, err = reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "$") {
			return nil, fmt.Errorf("expected bulk string, got %s", line)
		}

		length, err := strconv.Atoi(line[1:])

		if err != nil {
			return nil, err
		}

		// read bytes upto exact length
		buf := make([]byte, length)
		_, err = io.ReadFull(reader, buf)

		if err != nil {
			return nil, err
		}

		_, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		result = append(result, string(buf))
	}

	return result, nil
}
