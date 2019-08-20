package Global

import (
	"gopkg.in/mgo.v2"
	"net"
	"sync"
	"time"
)

var RoomCache Room          //cache满的时候新建一个room,将cache信息拷贝过去放进RoomMng字典中，清空cache room
var RoomMng map[int32]*Room //匹配满人添加一个房间，打完清除房间1-100
var RoomCacheMu sync.Mutex
var NextRoomID int32
var IDAddMu sync.Mutex
var ChanMap map[int32](chan []byte)
var NextUserID int32
var NextUserIDMu sync.Mutex
var RoomCollection *mgo.Collection
var UserCollection *mgo.Collection
var ClientMap map[net.Conn]ClientState

const RoomPeople int32 = 1
const FramesPerBag int32 = 3
const WaitMS time.Duration = 9
const DetailedLog string = "/Logs/"
