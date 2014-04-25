// Implementation of a MultiEchoServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"log"
	"strconv"
	"os"
	"bufio"
)

type Message struct {
	msg string
	conn net.Conn
}

type Clients struct {
	conn net.Conn
	outMesg chan string
}
	
type multiEchoServer struct {
	address string
	numClients int
	inMesg chan Message
	clients  []Clients
}

// New creates and returns (but does not start) a new MultiEchoServer.
func New() MultiEchoServer {
	// TODO: implement this!
	var MES MultiEchoServer
	serv := &multiEchoServer{address : "127.0.0.1",inMesg: make(chan Message)}
	go func(svr *multiEchoServer) {
		for {
			msg := <-svr.inMesg
			for _,cl:= range svr.clients {
				if cl.conn!=msg.conn {
					cl.outMesg<-msg.msg
				}
			}
		}
	}(serv)

	MES = serv
	
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
			defer c.Close()	
			outChan := make(chan string,6)
				
			mes.clients = append(mes.clients,Clients{conn: c,outMesg:outChan})

			go func(c net.Conn) {
				for{
					str:=<-outChan 
					_,err:=c.Write([]byte(str))
					checkError(err)
				}
			}(c)

			for {
				line,err:=bufio.NewReader(c).ReadBytes('\n')
				if err!= nil && err.Error()=="EOF" {
					break;
				}
				checkError(err)
				temp := Message{msg: string(line),conn: c}
				fmt.Println("got Msg: "+temp.msg)
				go func(m Message) { 
					mes.inMesg <- m	
				}(temp)
			}

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
		fmt.Fprintf(os.Stderr,"Error : %s \n",err.Error())
	}
}
