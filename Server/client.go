package main

import (
    "log"
    "net"
    "os"
    "../NetFrame"
)

func main() {
    client, err := net.Dial("tcp", "localhost:9700")
    if err != nil {
        log.Fatal("Client is dailing failed!")
        os.Exit(1)
    }

    //测试encode与deSerialize
    encode := NetFrame.NewEncode(8,1,1)
    encode.Write()
    //client.Write([]byte("i am client"))
    client.Write(encode.GetBytes())

    client.Close()
}