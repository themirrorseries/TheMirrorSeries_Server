package main

import (
	"../NetFrame"
	"log"
	"net"
	"os"
	//"../Handler"
	"../proto/dto"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

func main() {
	//client, err := net.Dial("tcp", "122.114.109.198:9700")
	//client, err := net.Dial("tcp", "localhost:9700")
	client, err := net.Dial("tcp", "mirror.murmur.top:9700")
	if err != nil {
		log.Fatal("Client is dailing failed!")
		os.Exit(1)
	}

	//测试encode与deSerialize
	anysend := DTO.MatchDTO{}
	anysend.Id = 1
	data, _ := proto.Marshal(&anysend)
	encode := NetFrame.NewEncode(int32(8+anysend.XXX_Size()), 2, 0)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)

	//client.Write([]byte("i am client"))
	client.Write(buffer.Bytes())
	//clinet.
	fmt.Println("正在匹配中...")
	var message []byte
	client.Read(message) //读到匹配成功消息

	//发送移动消息
	move := DTO.MoveDTO{}
	move.Roomid = 1
	move.Seat = 1
	move.X = 22.0
	move.Y = 33.0
	data2, _ := proto.Marshal(&anysend)
	encode2 := NetFrame.NewEncode(int32(8+move.XXX_Size()), 3, 0)
	encode2.Write()
	var buffer2 bytes.Buffer
	buffer2.Write(encode2.GetBytes())
	buffer2.Write(data2)
	client.Write(buffer2.Bytes())

	var message2 []byte
	client.Read(message2) //读到匹配成功消息
	fmt.Println("read server ok")
	/*
		var decode NetFrame.Decode

		decode.Read(message2)
		fmt.Println("decode ok")

		any := DTO.MoveDTO{}
		proto.Unmarshal(message2[decode.ReadPos:decode.Len+4], &any)
		fmt.Println("unmarshal ok")
		fmt.Println(decode.Len , " " , decode.Thetype , " " ,decode.Command , "",any.X, any.Y)
	*/
	var message3 []byte
	client.Read(message3) //读到匹配成功消息
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
	var ttt = 1000000
	for ttt < 0 {
		ttt--
	}

	client.Close()
}
