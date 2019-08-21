package main

import (
	"../Global"
	"../Handler"
	"github.com/sirupsen/logrus"
)

func main() {
	//初始化
	Global.RoomMng = make(map[int32]*Global.Room, 10)
	Global.RoomCache = *Global.NewRoom()
	Global.RoomCache.Clear()

	//Global.RoomCollection = Global.ConnecToRoom()
	//Global.UserCollection = Global.ConnecToUser()

	Global.NextRoomID = 1
	Global.NextUserID = 100000
	//Global.NextRoomID = Global.GetLastRoomID(Global.RoomCollection) //房间ID从1开始
	//Global.NextUserID = Global.GetLastUserID(Global.UserCollection) //用户ID从10万开始

	//初始化日志信息
	Global.DetailedLog.Log = Global.InitLog(logrus.InfoLevel, Global.DetailedLogPath)
	Global.ErrorLog.Log = Global.InitLog(logrus.ErrorLevel, Global.ErrorLogPath)
	Handler.Start()
}
