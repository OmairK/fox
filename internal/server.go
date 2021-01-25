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
	ConnID = 1
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
			delete(s.connections, id)
		}(ConnID)

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

func (ts *TCPServer) stop() {
	log.Printf("Server is stopping")
	close(ts.quit)
	ts.listner.Close()
}

//NewServer creates a new instance of Server
func NewServer(service string) *TCPServer {
	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal("Listener Error! ", err.Error())
	}

	srv := &TCPServer{
		listener:    listener,
		connections: map[uint]net.Conn{},
		quit:        make(chan struct{}),
	}
	go srv.listen()
	return srv
}

// func server() {
// 	arguments := os.Args
// 	var port string = ":8080"

// 	if len(arguments) == 1 {
// 		fmt.Println("Using port 8080")
// 	} else {
// 		port = arguments[1]
// 		fmt.Println("Using port ", port)
// 	}

// 	li, err := net.Listen("tcp", port)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer li.Close()

// 	conn, err := li.Accept()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, os.Interrupt)

// 	go func() {
// 		<-signalChan
// 		fmt.Println("Gracefully stopping the server")
// 		conn.Write([]byte("Server is stopping\n"))
// 		conn.Close()
// 	}()

// 	for {
// 		netData, err := bufio.NewReader(conn).ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if strings.TrimSpace(string(netData)) == "STOP" {
// 			fmt.Println("Exiting TCP server!")
// 			return
// 		}

// 		fmt.Print("-> ", string(netData))
// 		t := time.Now()
// 		myTime := t.Format(time.RFC3339) + "\n"
// 		conn.Write([]byte(myTime))
// 	}
// }
