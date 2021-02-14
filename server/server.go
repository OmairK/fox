package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/OmairK/fox/db"
)

func commandParser(action string, args int) ([]string, error) {
	action = strings.TrimSpace(action)
	res := strings.Split(action, " ")
	if len(res) == args {
		return res, nil
	}
	return nil, errors.New("Wrong number of arguments")

}

func performAction(action string, mdb *db.MemoryDB) string {
	if strings.ToUpper(action[:3]) == "GET" {
		key, err := commandParser(action[3:], 1)
		if err == nil {
			return mdb.Get(key[0])
		}
		return "Wrong number of arguments for 'get' command"
	} else if strings.ToUpper(action[:3]) == "SET" {
		args, err := commandParser(action[3:], 2)
		if err == nil {
			return mdb.Set(args[0], args[1])
		}
		return "Wrong number of arguments for 'set' command"
	}
	return "Invalid Command"
}

// TCPServer is the multi threaded tcp server
type TCPServer struct {
	Listner     net.Listener
	Connections map[uint]net.Conn
	Quit        chan struct{}
	Database    *db.MemoryDB
}

func (ts *TCPServer) listen() {
	var ConnID uint
	ConnID = 0
	fmt.Println("Waiting for connections")
	for {
		conn, err := ts.Listner.Accept()
		if err != nil {
			if _, ok := <-ts.Quit; ok {
				log.Println("New Connection error!", err.Error())
				continue
			} else {
				ts.commit()
				break
			}
		}
		ts.Connections[ConnID] = conn
		go func(id uint) {
			log.Printf("Connection with ID %d joined!", id)
			ts.handleConnection(conn)
			log.Printf("Connection with ID %d leaving! ", id)
			conn.Close()
			delete(ts.Connections, id)
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
		fmt.Print(netData)
		result := performAction(netData, ts.Database) + "\n"
		conn.Write([]byte(result))
	}

}

// Stop handles the stopping of the TCPServer
func (ts *TCPServer) Stop() {
	log.Printf("Server is stopping")
	ts.broadcast("I dont feel so great Mr Stark...")
	close(ts.Quit)
	ts.Listner.Close()
}

func (ts *TCPServer) broadcast(message string) {
	message = message + "\n"
	for _, conn := range ts.Connections {
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
		Listner:     listner,
		Connections: map[uint]net.Conn{},
		Quit:        make(chan struct{}),
		Database: &db.MemoryDB{
			Name: "new",
			KeyV: make(map[string]string),
		},
	}
	go srv.listen()
	return srv
}
