package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, string('*')) {
		return nil, fmt.Errorf("expected a RESP array, got %s\n", line)
	}

	count, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to read RESP array length %s\n", err.Error())
	}

	result := make([]string, 0, count)

	for range count {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, string('$')) {
			return nil, fmt.Errorf("expected a bulk string, got %s\n", line)
		}

		length, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("failed to read the bulk string size %s", err.Error())
		}

		buffer := make([]byte, length)

		_, err = io.ReadFull(reader, buffer)

		if err != nil {
			return nil, err
		}

		// cleaning \r\n after the string
		_, err = reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		result = append(result, string(buffer))
	}
	return result, nil
}
