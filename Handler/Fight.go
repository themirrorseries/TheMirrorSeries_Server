package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	//log "github.com/sirupsen/logrus"
	//"fmt"
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
	case int32(DTO.FightTypes_DEATH_CREQ): //客户端死亡
		fight.death()
		break
	case int32(DTO.FightTypes_GAME_OVER_CREQ): //游戏结束
		fight.gameOver()
		break
	case int32(DTO.FightTypes_LOAD_UP_CREQ):
		fight.loadUp()
		break
	case int32(DTO.FightTypes_LAEVE_CREQ):
		fight.leaveRoom()
		break
	default:
		break
	}
}

func (fight *Fight) move() {

	move := DTO.ClientMoveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &move)
	//if Global.RoomMng[move.Roomid]
	_, ok := Global.RoomMng[move.Roomid]
	if ok {
		Global.RoomMng[move.Roomid].InsertMsg(&move)
	}
}

func (fight *Fight) death() {
	death := DTO.FightLeaveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &death)
	//fmt.Println("死亡", death.Seat)
	Global.RoomMng[death.Roomid].PlayerDeath(death.Seat)
}

func (fight *Fight) gameOver() {
	gameOver := DTO.FightLeaveDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &gameOver)
	//fmt.Println("游戏结束", gameOver.Seat)
	Global.RoomMng[gameOver.Roomid].GameOver(gameOver.Seat)
}
func (fight *Fight) loadUp() {
	load := DTO.FightLoadDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &load)
	Global.RoomMng[load.Roomid].AddLoadPeople(load.Seat)
}

func (fight *Fight) leaveRoom() {
	leave := DTO.FightLoadDTO{}
	proto.Unmarshal(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], &leave)
	//fmt.Println("离开", leave.Seat)
	Global.RoomMng[leave.Roomid].PlayerLeave(leave.Seat)
}
