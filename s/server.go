package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// サーバーを起動し、クライアントからの接続を待機する
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server started. Waiting for connections...")

	for {
		// クライアントからの接続を受け入れる
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		// クライアントとの通信を処理するゴルーチンを開始する
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// クライアントからのデータを受信する
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err.Error())
		return
	}

	// 受信したデータを文字列に変換し、猫語に変換する
	message := string(buffer[:n])
	catMessage := convertToCatLanguage(message)

	// 猫語に変換したデータをクライアントに送信する
	_, err = conn.Write([]byte(catMessage))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}

	fmt.Println("Message sent to client:", catMessage)
}

func convertToCatLanguage(message string) string {
	// 猫語に変換するロジックを実装する
	// 2ビットずつ取り出して猫語に変換する例を示します

	var builder strings.Builder
	for _, char := range message {
		// 文字コードをバイトスライスに変換
		bytes := []byte(string(char))

		// バイトスライスの各バイトを2ビットずつ取り出して猫語に変換
		for _, b := range bytes {
			// バイトの上位6ビットを無視
			b = b & 0x03

			// 2ビットごとに変換
			switch b {
			case 0x00:
				builder.WriteString("ニャン")
			case 0x01:
				builder.WriteString("ミャン")
			case 0x02:
				builder.WriteString("ミャオン")
			case 0x03:
				builder.WriteString("ニャーン")
			}
		}
	}

	return builder.String()
}
