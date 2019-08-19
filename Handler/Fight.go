package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	//log "github.com/sirupsen/logrus"
)

type Fight struct {
	data *HandlerData
}

func (fight *Fight) ReceiveMessage() {
	switch fight.data.command {
	case int32(DTO.FightTypes_MOVE_CREQ):
		//client move
		fight.move()
	}
}

func (fight *Fight) move() {

	move := DTO.ClientMoveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &move)

	Global.RoomMng[move.Roomid].InsertMsg(&move)
}
