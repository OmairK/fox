package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// TCPServer is the multi threaded tcp server
type TCPServer struct {
	listner     net.Listener
	connections map[uint]net.Conn
	quit        chan struct{}
}

func (ts *TCPServer) listen() {
	var ConnID uint
	ConnID = 0
	fmt.Println("Waiting for connections")
	for {
		conn, err := ts.listner.Accept()
		if err != nil {
			if _, ok := <-ts.quit; ok {
				log.Println("New Connection error!", err.Error())
				continue
			} else {
				ts.commit()
				break
			}
		}
		ts.connections[ConnID] = conn
		go func(id uint) {
			log.Printf("Connection with ID %d joined!", id)
			ts.handleConnection(conn)
			log.Printf("Connection with ID %d leaving! ", id)
			conn.Close()
			delete(ts.connections, id)
		}(ConnID)
		ConnID++

	}
}

func (ts *TCPServer) commit() {
	fmt.Println("Commiting the changes")
}

func (ts *TCPServer) handleConnection(conn net.Conn) {
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))
	}

}

// Stop handles the stopping of the TCPServer
func (ts *TCPServer) Stop() {
	log.Printf("Server is stopping")
	ts.broadcast("I dont feel so great Mr Stark...")
	close(ts.quit)
	ts.listner.Close()
}

func (ts *TCPServer) broadcast(message string) {
	message = message + "\n"
	for _, conn := range ts.connections {
		conn.Write([]byte(message))
	}

}

//NewServer creates a new instance of Server
func NewServer(service string) *TCPServer {
	listner, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal("Listener Error! ", err.Error())
	}

	srv := &TCPServer{
		listner:     listner,
		connections: map[uint]net.Conn{},
		quit:        make(chan struct{}),
	}
	go srv.listen()
	return srv
}
