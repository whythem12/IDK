package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"time"
)
var (
	client    []Clients
	curClient int
	curConn   net.Conn
	reader    *bufio.Reader
	timeout   = 500 * time.Millisecond
)

type Clients struct {
	Client    net.Conn 
	IPAddress string
}

var connection net.Conn
func sendMessage() {
	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		textInput, _ := reader.ReadString('\n')
		textInput = strings.Replace(textInput, "\r", "", -1)
		textInput = strings.Replace(textInput, "\n", "", -1)
		if len(textInput) != 0 {
			connection.Write([]byte(textInput))
		}
	}
}
func main() {
	conn, err := net.Dial("tcp", ":500")
	if err != nil {
		fmt.Println(err)
	}
	connection = conn
	go checkMessage(); go sendMessage()

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		data := string(buf[:size])
		fmt.Println("Received From Server: " + data)
	
	}

}

func checkMessage() {
	for {
		for i, v := range client {
			v.Client.SetReadDeadline(time.Now().Add(timeout))
			buf := make([]byte, 1024)
			size, _ := v.Client.Read(buf)
			data := string(buf[:size])
			if size != 0 {
				fmt.Println("We Got This From Client " + strconv.Itoa(i) + ": " + data)
			}
		}
	}
}
