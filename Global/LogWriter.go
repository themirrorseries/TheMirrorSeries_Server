package Global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

//请求匹配/手动退出匹配/进入战斗房间/死亡或战斗结束退出战斗房间/手动强制退出战斗房间/战斗死亡/匹配中断开连接/战斗中断开连接/（其他）断开连接/
type MyLog struct {
	Log *logrus.Logger
}
type LogEventTypes int32

const (
	Login_IN            LogEventTypes = 0
	Match_IN            LogEventTypes = 1
	Match_OUT           LogEventTypes = 2
	Fight_IN            LogEventTypes = 3
	Fight_OUT_NORMAL    LogEventTypes = 4
	Fight_OUT_FORCE     LogEventTypes = 5
	Fight_Death         LogEventTypes = 6
	Client_OUT_Matching LogEventTypes = 7
	Client_OUT_Fighting LogEventTypes = 8
	Client_OUT_Other    LogEventTypes = 9
)

var EventTypes_name = map[LogEventTypes]string{
	0: "登录",
	1: "请求匹配",
	2: "手动退出匹配",
	3: "进入战斗房间",
	4: "死亡或战斗结束退出战斗房间",
	5: "手动强制退出战斗房间",
	6: "战斗死亡",
	7: "匹配中断开连接",
	8: "战斗中断开连接",
	9: "其他情况断开连接",
}

func InitLog(logLevel logrus.Level, filepath string) *logrus.Logger {
	var log = logrus.New()
	//log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                      //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true     // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output
	log.Level = logLevel

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	return log
}

//玩家id	请求匹配/手动退出匹配/进入战斗房间/死亡或战斗结束退出战斗房间/手动强制退出战斗房间/战斗死亡/匹配中断开连接/战斗中断开连接/（其他）断开连接/
func (myLog *MyLog) Detailed(playerid int32, eventTpye LogEventTypes) {
	myLog.Log.Info(fmt.Sprintf("---玩家id:%d%s---", playerid, EventTypes_name[eventTpye]))
}

//手动强制退出战斗房间/匹配中断开连接/战斗中断开连接/（其他）断开连接/
func (myLog *MyLog) Warn(playerid int32, eventTpye LogEventTypes) {
	myLog.Log.Warn(fmt.Sprintf("---玩家id:%d%s---", playerid, EventTypes_name[eventTpye]))
}
