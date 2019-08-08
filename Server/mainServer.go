package main

import (	"../Global"
"../Handler"
)
func main() {
	//初始化
	Global.RoomCache = *Global.NewRoom()
	Global.RoomCache.Clear()
	Global.NextRoomID = 1
	Global.NextUserID = 100000
	Handler.Start()
}
