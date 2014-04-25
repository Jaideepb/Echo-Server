// Implementation of a MultiEchoServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"log"
	"strconv"
	"io/ioutil"
	"os"
)

type Message struct {
	msg string
	conn net.Conn
}
	
type multiEchoServer struct {
	// TODO: implement this!
	address string
	numClients int
	inMesg chan Message
	clients  []net.Conn
}

// New creates and returns (but does not start) a new MultiEchoServer.
func New() MultiEchoServer {
	// TODO: implement this!
	var MES MultiEchoServer
	serv := multiEchoServer{address : "127.0.0.1",inMesg: make(chan Message)}
	go func(svr multiEchoServer) {
		for {
			msg := <-svr.inMesg
			for _,cl:= range svr.clients {
				if cl!=msg.conn {
					_,err:=cl.Write([]byte(msg.msg))
					checkError(err)
				}
			}
		}
	}(serv)

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
		conn,err := l.Accept()
		if(err!=nil) {
			log.Fatal(err)
			return err
		}
		go func(c net.Conn) {
			mes.clients = append(mes.clients,c)
			go func(c net.Conn) {
				for {
					res,err:= ioutil.ReadAll(c)
					checkError(err)
					temp := Message{msg: string(res),conn: c}
					go func(m Message) { 
						mes.inMesg <- m	
					}(temp)
				}
			}(c)
		}(conn)

		mes.numClients++;
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

func checkError(err error) {
	if(err!=nil) {
		fmt.Fprintf(os.Stderr,"Error : %s",err.Error())
	}
}
