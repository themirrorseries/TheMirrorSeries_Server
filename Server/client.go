package main

import (
	"../NetFrame"
	"log"
	"net"
	"os"
	//"../Handler"
	"fmt"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	"time"
	"bytes"
)

func main() {
	client, err := net.Dial("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("Client is dailing failed!")
		os.Exit(1)
	}

	//测试encode与deSerialize
	anysend :=DTO.AnyDTO{}
	anysend.Code = 1
	data, _ :=proto.Marshal(&anysend)
	encode := NetFrame.NewEncode(int32(12+anysend.XXX_Size()), 0, 0)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)


	//client.Write([]byte("i am client"))
	client.Write(buffer.Bytes())
	//clinet.
	var message []byte
	client.Read(message)
	var decode NetFrame.Decode
	decode.Read(message)
	any := DTO.AnyDTO{}
	proto.Unmarshal(message[decode.ReadPos:decode.Len], &any)
	fmt.Println(decode.Len , " " , decode.Thetype , " " ,decode.Command , "",any.Code )

	timeTicker := time.NewTicker(time.Second * 10)
	i := 0
	for {
		if i > 5 {
			break
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		i++
		<-timeTicker.C

	}
	// 清理计时器
	timeTicker.Stop()
	client.Close()
}
