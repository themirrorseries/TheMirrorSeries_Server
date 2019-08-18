package Handler

import (
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1024)
	//var err error=!nil

	for {
		log.Error("开始读数据")
		len, err := client.Read(message)
		log.Error("读数据完成")
		//todo it doesn't work, unless we add heart bag
		if err != nil {
			//return client.Close()
			log.Error("client out", client.Close())
			break
		}
		if len > 0 {
			//log.Println(string(message[0:len]))
			Handler(message[0:len], client)
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
