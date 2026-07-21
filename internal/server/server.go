package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/saicharith2012/goredis/internal/command"
	"github.com/saicharith2012/goredis/internal/resp"
	"github.com/saicharith2012/goredis/internal/store"
)

type Server struct {
	port  string
	store *store.Store
}

func New(port string) *Server {
	store := store.New()
	return &Server{port, store}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.port)

	if err != nil {
		return fmt.Errorf("error while listening for connections %s", err.Error())
	}

	defer listener.Close()

	fmt.Printf("goredis running on port %s...\n", s.port[1:])

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("error trying to connect: %s\n", err)
			continue
		}

		fmt.Println("client connected:", conn.RemoteAddr())

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		tokens, err := resp.Parse(reader)

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("client closed the connection:", conn.RemoteAddr())
			} else {
				fmt.Println("read error:", err.Error())
				fmt.Println("closed the client connection", conn.RemoteAddr())
			}
			return
		}

		var response string

		if len(tokens) > 0 {
			fmt.Printf("Received chunk %#v from client %s\n", tokens, conn.RemoteAddr())
		}

		response = command.HandleCommand(tokens, s.store)
		_, err = writer.Write([]byte(response))

		if err != nil {
			fmt.Println("failed to write:", err.Error())
			return
		}

		err = writer.Flush()

		if err != nil {
			fmt.Println("unable to write to connection:", err.Error())
			fmt.Println("closed the client connection", conn.RemoteAddr())
			return
		}

	}
}
