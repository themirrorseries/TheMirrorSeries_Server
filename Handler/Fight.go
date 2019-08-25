package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	//log "github.com/sirupsen/logrus"
	"fmt"
)

type Fight struct {
	data *HandlerData
}

//处理战斗过程中通信 包括移动（放技能）/场景加载完成
func (fight *Fight) ReceiveMessage() {
	switch fight.data.command {
	case int32(DTO.FightTypes_MOVE_CREQ):
		fight.move()
		break
	case int32(DTO.FightTypes_LIVEROOM_CREQ): //客户端离开
		fight.leaveRoom()
		break
	case int32(DTO.FightTypes_DEATH_CREQ): //客户端死亡
		fight.death()
		break
	case int32(DTO.FightTypes_WIN_CREQ): //客户端胜利
		break
	case int32(DTO.FightTypes_LOAD_UP_CREQ):
		fight.loadUp()
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
func (fight *Fight) loadUp() {
	fmt.Println("client load ")
	load := DTO.FightLoadDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &load)
	Global.RoomMng[load.Roomid].AddLoadPeople(load.Seat)
}
