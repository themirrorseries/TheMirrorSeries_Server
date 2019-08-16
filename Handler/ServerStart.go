package Handler

import (
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1024)

	for {
		//todo it doesn't work, unless we add heart bag
		if client == nil {
			//return client.Close()
			log.Error("client out", client.Close())
			break
		}
		len, _ := client.Read(message)
		if len > 0 {
			//log.Println(string(message[0:len]))
			DeSerizalize(message[0:len], client)
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
