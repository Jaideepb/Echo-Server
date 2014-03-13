// Implementation of a MultiEchoServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"log"
	"strconv"
)
	
type multiEchoServer struct {
	// TODO: implement this!
	address string
	numClients int
}

// New creates and returns (but does not start) a new MultiEchoServer.
func New() MultiEchoServer {
	// TODO: implement this!
	var MES MultiEchoServer
	serv := multiEchoServer{address : "1.0.0.0"}
	MES = &serv
	
	if(MES!=nil) {
		return MES
	} else {
		return nil
	}
}

func (mes *multiEchoServer) Error() string {
	return fmt.Sprintf("Unable to start the server on address: %s",mes.address)
} 

func (mes *multiEchoServer) Start(port int) error {
	// TODO: implement this!
	l,err := net.Listen("tcp",mes.address + ":" + strconv.Itoa(port))
	if(err!=nil) {
		return mes
	}
	
	for {
		_,err := l.Accept()
		if(err!=nil) {
			log.Fatal(err)
		}
	}
	return nil
}

func (mes *multiEchoServer) Close() {
	// TODO: implement this!
}

func (mes *multiEchoServer) Count() int {
	// TODO: implement this!
	return -1
}

// TODO: add additional methods/functions below!
