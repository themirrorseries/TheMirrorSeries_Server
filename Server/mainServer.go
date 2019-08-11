package main

import (	"../Global"
"../Handler"

)
func main() {
	//初始化
	Global.RoomMng = make(map[int32]*Global.Room, 10)
	Global.RoomCache = *Global.NewRoom()
	Global.RoomCache.Clear()
	Global.NextRoomID = 1
	Global.NextUserID = 100000
	Handler.Start()
}
