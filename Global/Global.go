package Global

import ("sync"
"../Handler"
	)

var RoomCache		Handler.Room	//cache满的时候新建一个room,将cache信息拷贝过去放进RoomMng字典中，清空cache room
var RoomMng			map[int32]Handler.Room	//匹配满人添加一个房间，打完清除房间
var RoomCacheMu		sync.Mutex
var	NextRoomID		int32
var IDAddMu			sync.Mutex