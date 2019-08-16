package main

import (
	"../NetFrame"
	"../proto/dto"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

func main() {
	//client, err := net.Dial("tcp", "122.114.109.198:9700")
	client, err := net.Dial("tcp", "localhost:9700")
	//client, err := net.Dial("tcp", "mirror.murmur.top:9700")
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
	client.Write(buffer.Bytes())

	fmt.Println("正在匹配中...")
	/*
		var message4 []byte
		message4 = make([]byte, 1024)
		client.Read(message4) //读到匹配返回消息
		//------------取消匹配测试
		removeMatch := DTO.MatchRtnDTO{}
		removeMatch.Id = 1
		dataRemove, _ := proto.Marshal(&removeMatch)
		encodeRemove := NetFrame.NewEncode(int32(8+removeMatch.XXX_Size()), 2, 2)
		encodeRemove.Write()
		var bufferRemove bytes.Buffer
		bufferRemove.Write(encodeRemove.GetBytes())
		bufferRemove.Write(dataRemove)
		client.Write(bufferRemove.Bytes())
		fmt.Println("取消匹配...")
		var message []byte
		message = make([]byte, 1024)
		client.Read(message) //读到取消匹配消息
	*/
	//发送移动消息
	move := DTO.ClientMoveDTO{}
	Msg := make([]DTO.FrameInfo, 1)
	move.Msg = make([]*DTO.FrameInfo, 1)
	move.Msg[0] = &Msg[0]
	move.Roomid = 1
	move.Seat = 2
	move.Bagid = 1
	move.Msg[0].Frame = 1

	move.Msg[0].Move = new(DTO.DeltaDirection)
	move.Msg[0].Move.X = 11
	move.Msg[0].Move.Y = 22
	move.Msg[0].Move.DeltaTime = 1

	data2, _ := proto.Marshal(&move)

	encode2 := NetFrame.NewEncode(int32(8+move.XXX_Size()), 3, 0)
	encode2.Write()
	var buffer2 bytes.Buffer
	buffer2.Write(encode2.GetBytes())
	buffer2.Write(data2)
	client.Write(buffer2.Bytes())

	var message2 []byte
	message2 = make([]byte, 1024)
	len, _ := client.Read(message2) //读到移动消息
	//client.Read(message2) //读到移动消息
	fmt.Println("read server ok")
	var decode NetFrame.Decode

	decode.Read(message2[0:len])
	fmt.Println("decode ok")

	any := DTO.ServerMoveDTO{}
	proto.Unmarshal(message2[decode.ReadPos:decode.Len+4], &any)
	fmt.Println("unmarshal ok")
	fmt.Println(decode.Len, " ", decode.Thetype, " ", decode.Command, "", any.Bagid, any.ClientInfo[0].Msg[0].Move.X)

	var message3 []byte
	message3 = make([]byte, 1024)
	client.Read(message3) //读到匹配成功消息
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
