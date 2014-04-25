package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
//	"io"
	"time"
)

const (
	defaultHost = "localhost"
	defaultPort = 9999
)

// To test your server implementation, you might find it helpful to implement a
// simple 'client runner' program. The program could be very simple, as long as
// it is able to connect with and send messages to your server and is able to
// read and print out the server's echoed response to standard output. Whether or
// not you add any code to this file will not affect your grade.
func main()  {
	if len(os.Args) !=2 {
		fmt.Fprintf(os.Stderr,"Usage wrong")
	}	
	
	addr:= os.Args[1]
	conn,err:= net.Dial("tcp",addr)
	checkError(err)

	sync :=make(chan int)

	go func(c net.Conn) {
		
		for {
			bio:= bufio.NewReader(os.Stdin)
			line,_,err:=bio.ReadLine()
			mesg:= string(line)
			mesg+="\n"

			checkError(err)
			if(err!=nil) {
				break
			}
			
			c.Write([]byte(mesg))
		}
		sync<-1
	}(conn)
	go ReadFromSer(conn)
	
	<-sync
	os.Exit(0)
}

func ReadFromSer(c net.Conn) {
	for{		
		time.Sleep(10*time.Second)
		var b []byte
		b=make([]byte,100)
		c.Read(b)
		fmt.Printf(string(b))
	}
}

func checkError(err error) {
	if(err!=nil) {
		fmt.Fprintf(os.Stderr,"Error: %s",err.Error())
		os.Exit(1)
	}
}
