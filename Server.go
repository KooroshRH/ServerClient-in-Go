package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting server...")
	server, _ := net.Listen("tcp", ":8080")
	connect, _ := server.Accept()
	fmt.Println("Connected!!")
	for {
		operation, _ := bufio.NewReader(connect).ReadString('\n')
		fmt.Println("Received")
		if strings.Contains(operation, "+") {
			operators := strings.Split(operation, "+")
			operators[1] = strings.TrimSpace(operators[1])
			first, _ := strconv.Atoi(operators[0])
			sec, _ := strconv.Atoi(operators[1])
			sum := first + sec
			result := strconv.Itoa(sum)
			connect.Write([]byte(result + "\n"))
			fmt.Println("Sent!!")
		} else if strings.Contains(operation, "*") {
			operators := strings.Split(operation, "*")
			operators[1] = strings.TrimSpace(operators[1])
			first, _ := strconv.Atoi(operators[0])
			sec, _ := strconv.Atoi(operators[1])
			sum := first * sec
			result := strconv.Itoa(sum)
			connect.Write([]byte(result + "\n"))
			fmt.Println("Sent!!")
		} else if strings.Contains(operation, "-") {
			operators := strings.Split(operation, "-")
			operators[1] = strings.TrimSpace(operators[1])
			first, _ := strconv.Atoi(operators[0])
			sec, _ := strconv.Atoi(operators[1])
			sum := first - sec
			result := strconv.Itoa(sum)
			connect.Write([]byte(result + "\n"))
			fmt.Println("Sent!!")
		}
	}
}
