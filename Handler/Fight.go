package Handler

import (
	"../proto/dto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"../Global"
)

type Fight struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     net.Conn
}

func NewFight(c, start, end int32, msg []byte, _client net.Conn) *Fight {
	fight := &Fight{
		command:    c,
		bytesStart: start,
		bytesEnd:   end,
		messages:   msg,
		client:     _client,
	}
	return fight
}

func (fight *Fight) ReveiveMessage() {
	switch fight.command {
	case 0:
		//client move
		fight.move()
	}
}

func (fight *Fight) move() {
	fmt.Println("move start")

	move := DTO.MoveDTO{}
	proto.Unmarshal(fight.messages[fight.bytesStart:fight.bytesEnd], &move)
	//Global.RoomMng[move.Roomid].RoomBroad(&move, move.Seat)
	Global.RoomMng[move.Roomid].RoomBroad(&move)
}
