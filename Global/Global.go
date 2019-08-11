package Global

import (
	"../proto/dto"
	"sync"
)

//自定义消息帧
type FrameMessage struct {
	FrameNum int32
	Message  []MessageNode
}

type MessageNode struct {
	move     DTO.MoveDTO
	NextNode *MessageNode
}

func NewMsgNode(_move DTO.MoveDTO) *MessageNode {
	msgNode := &MessageNode{
		move:     _move,
		NextNode: nil,
	}
	return msgNode
}

var RoomCache Room          //cache满的时候新建一个room,将cache信息拷贝过去放进RoomMng字典中，清空cache room
var RoomMng map[int32]*Room //匹配满人添加一个房间，打完清除房间
var RoomCacheMu sync.Mutex
var NextRoomID int32
var IDAddMu sync.Mutex
var ChanMap map[int32](chan []byte)
var NextUserID int32
var NextUserIDMu sync.Mutex

const RoomPeople int32 = 2
