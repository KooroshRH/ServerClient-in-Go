package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connect, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println("Connected!!")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your operations: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connect, text+"\n")
		fmt.Println("Sent!!")
		result, _ := bufio.NewReader(connect).ReadString('\n')
		fmt.Println("Result: " + result)
	}
}
