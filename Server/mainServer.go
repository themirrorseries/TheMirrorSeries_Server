package main

import (
	"../Global"
	"../Handler"
	"github.com/sirupsen/logrus"
	//"net/http"
	//_ "net/http/pprof"
)

func main() {
	//初始化
	Global.RoomMng = make(map[int32]*Global.Room, 10)
	Global.RoomCache = *Global.NewRoom()
	Global.RoomCache.Clear()

	if Global.UseMongo {
		Global.RoomCollection = Global.ConnecToRoom()
		Global.UserCollection = Global.ConnecToUser()
		Global.NextRoomID = Global.GetLastRoomID(Global.RoomCollection) //房间ID从1开始
		Global.NextUserID = Global.GetLastUserID(Global.UserCollection) //用户ID从10万开始
	} else {
		Global.NextRoomID = 1
		Global.NextUserID = 100000
	}

	//初始化日志信息
	Global.DetailedLog.Log = Global.InitLog(logrus.InfoLevel, Global.DetailedLogPath)
	Global.ErrorLog.Log = Global.InitLog(logrus.ErrorLevel, Global.ErrorLogPath)
	//go func() {
	//	http.ListenAndServe(":6060", nil)
	//}()
	Handler.Start()
}
