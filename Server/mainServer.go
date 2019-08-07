package main

import ("../NetFrame"
	"../Global"
)
func main() {
	//初始化
	Global.RoomCache.Clear()
	Global.NextRoomID = 1
	NetFrame.Start()

}
