package Handler

import (
	"../Global"
	//"bytes"
	//"encoding/binary"
	//"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	clientState := Global.NewClientState(client)
	var message []byte
	message = make([]byte, 1024)

	for {
		len, err := client.Read(message)
		if err != nil {
			//客户端断开处理——匹配中
			if clientState.IsMatch {
				clientState.MatchOut()
				Global.DetailedLog.Warn(clientState.PlayerID, Global.Client_OUT_Matching)
			}
			//客户端断开处理——在战斗中
			if clientState.IsFight {
				clientState.FightOut()
				Global.DetailedLog.Warn(clientState.PlayerID, Global.Client_OUT_Fighting)
			}
			//log.Error("client out", client.Close())
			break
		}
		/*
			readPos := int32(0)
			var msglen int32
			for {
				if len > 0 {
					binary.Read(bytes.NewBuffer(message[readPos:]), binary.LittleEndian, &msglen)
					msglen += 4
					if len > int(msglen) {
						fmt.Println("发生粘包：", len, msglen)
					}
					Handler(message[readPos:readPos+msglen], clientState)
					readPos += msglen
					len -= int(msglen)
				} else {
					break
				}
			}*/
		if len > 0 {
			//Handler(message[readPos:readPos+msglen], clientState)
			Handler(message[0:len], clientState)
		}
	}
	return nil
}

func Start() {
	server, err := net.Listen("tcp", "0.0.0.0:9700")
	if err != nil {
		Global.ErrorLog.Log.Errorln("start server failed!\n")
		os.Exit(1)
	}
	defer server.Close()
	log.Println("server is running...")
	Global.DetailedLog.Log.Info("服务端开启...")
	for {
		client, err := server.Accept()
		if err != nil {
			Global.ErrorLog.Log.Errorln("Accept error\n")
			continue
		}

		//log.Println("the client is connectted...")
		Global.DetailedLog.Log.Info("一个新的客户端连接...")
		go recvMessage(client)
	}
}
