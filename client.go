package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// サーバーに接続する
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// クライアントからの入力を受け取る
	fmt.Print("Enter a message: ")
	var input string
	fmt.Scanln(&input)

	// サーバーに入力を送信する
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("Error sending data to server:", err.Error())
		return
	}

	// サーバーからの応答を受信する
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data from server:", err.Error())
		return
	}

	// 受信した応答を表示する
	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}

