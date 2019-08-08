package main

import (
	"../NetFrame"
	"log"
	"net"
	"os"
	//"../Handler"
)

func main() {
	client, err := net.Dial("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("Client is dailing failed!")
		os.Exit(1)
	}

	//测试encode与deSerialize
	encode := NetFrame.NewEncode(8, 2, 0)
	encode.Write()
	//client.Write([]byte("i am client"))
	client.Write(encode.GetBytes())
	//clinet.
	client.Close()
}
