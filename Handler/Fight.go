package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	//log "github.com/sirupsen/logrus"
	"../NetFrame"
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
	aesDec, err := NetFrame.AesDecrypt(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], NetFrame.MyKey)
	if err != nil {
		Global.ErrorLog.Log.Errorln(err, "客户端解码出错！")
	}
	death := DTO.FightLeaveDTO{}
	proto.Unmarshal(aesDec, &death)
	Global.RoomMng[death.Roomid].PlayerDeath(death.Seat)
}

func (fight *Fight) leaveRoom() {
	aesDec, err := NetFrame.AesDecrypt(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], NetFrame.MyKey)
	if err != nil {
		Global.ErrorLog.Log.Errorln(err, "客户端解码出错！")
	}
	leave := DTO.FightLeaveDTO{}
	proto.Unmarshal(aesDec, &leave)
	Global.RoomMng[leave.Roomid].PlayerLeave(leave.Seat)
}

func (fight *Fight) winGame() {
	//房间号  游戏信息
}
func (fight *Fight) loadUp() {
	fmt.Println("client load ")
	aesDec, err := NetFrame.AesDecrypt(fight.data.messages[fight.data.bytesStart:fight.data.bytesEnd], NetFrame.MyKey)
	if err != nil {
		Global.ErrorLog.Log.Errorln(err, "客户端解码出错！")
	}
	load := DTO.FightLoadDTO{}
	proto.Unmarshal(aesDec, &load)
	Global.RoomMng[load.Roomid].AddLoadPeople(load.Seat)
	fmt.Println(load.Seat)
}
