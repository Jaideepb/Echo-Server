package main

import (
	"fmt"
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

	//_,err= conn.Write([]byte("Hello\r\n"))
	//checkError(err)
	
		
	
	os.Exit(0)
}

func checkError(err error) {
	if(err!=nil) {
		fmt.Fprintf(os.Stderr,"Error: %s",err.Error())
	}
}
