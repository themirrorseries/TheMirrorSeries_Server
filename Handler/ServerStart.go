package Handler

import (
	"../Global"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	clientState := Global.NewClientState(client)
	var message []byte
	message = make([]byte, 1024)
	//var err error=!nil

	for {
		len, err := client.Read(message)
		if err != nil {
			//客户端断开处理	匹配中 房间中
			if clientState.IsMatch {
				clientState.MatchOut()
			}
			if clientState.IsFight {
				clientState.FightOut()
			}
			log.Error("client out", client.Close())
			break
		}
		if len > 0 {
			//log.Println(string(message[0:len]))
			Handler(message[0:len], clientState)
		}
	}
	return nil
}

func Start() {
	server, err := net.Listen("tcp", "0.0.0.0:9700")
	if err != nil {
		log.Fatal("start server failed!\n")
		os.Exit(1)
	}
	defer server.Close()
	log.Println("server is running...")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error\n")
			continue
		}

		log.Println("the client is connectted...")
		go recvMessage(client)
	}
}
