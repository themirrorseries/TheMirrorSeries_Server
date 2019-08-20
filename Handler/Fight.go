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
		break
	case 3: //客户端离开
		fight.leaveRoom()
		break
	case 4: //客户端死亡
		fight.death()
		break
	case 5: //客户端胜利
		break
	default:
		break
	}
}

func (fight *Fight) move() {

	move := DTO.ClientMoveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &move)

	Global.RoomMng[move.Roomid].InsertMsg(&move)
}

func (fight *Fight) death() {
	//房间号 Seat号
	death := DTO.FightLeaveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &death)
	Global.RoomMng[death.Roomid].PlayerDeath(death.Seat)
}

func (fight *Fight) leaveRoom() {
	//房间号 Seat号
	death := DTO.FightLeaveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &death)
	Global.RoomMng[death.Roomid].PlayerLeave(death.Seat)
}

func (fight *Fight) winGame() {
	//房间号  游戏信息
}
