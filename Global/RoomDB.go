package Global

import (
	"../proto/dto"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func ConnecToRoom() *mgo.Collection {
	session, err := mgo.Dial(MongoURL)
	if err != nil {
		ErrorLog.Log.Errorln("Room数据库连接出错！")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DataBase).C(RoomDB)
	return c
}

func GetLastRoomID(c *mgo.Collection) int32 {
	count, _ := c.Count()
	if count == 0 {
		return 1
	} else {
		fmt.Println(count)
		var a MongoDBRoomMng
		c.Find(nil).Sort("-roomid").One(&a)
		return a.RoomID + 1
	}
}

//新开一个房间就新建一个Room数据
func AddRoom(c *mgo.Collection, roomid int32, match *DTO.MatchSuccessDTO) {

	//player信息房间匹配好就传进来建立表
	players := make([]Player, RoomPeople)
	//frame := make([]Frame, RoomPeople*FramesPerBag)

	for i := int32(0); i < RoomPeople; i++ {
		players[i] = Player{match.Players[i].Playerid, match.Players[i].Name, match.Players[i].Roleid, match.Players[i].Seat}
	}
	room := MongoDBRoom{players, []Frame{}}
	roomMng := MongoDBRoomMng{"1", roomid, room}
	c.Insert(&roomMng)
	DetailedLog.Log.Info("新建房间数据库，房间号为---", roomMng.RoomID)
}

//每次插入RoomPeople*FramesPerBag帧
func AddFrame(c *mgo.Collection, roomid int32, t chan *DTO.ServerMoveDTO) {
	timer := time.NewTimer(WaitMS * time.Second)
	flag := false
	for {
		timer = time.NewTimer(10 * time.Second)
		go func() {
			select {
			case <-timer.C:
				{
					close(t)
					flag = true
					return
				}
			}
		}()
		send := <-t
		timer.Stop()
		roomFrames := make([]Frame, RoomPeople*FramesPerBag)

		for i := int32(0); i < RoomPeople; i++ {
			if send.ClientInfo[i].Seat == -1 {
				break
			}
			for j := int32(0); j < FramesPerBag; j++ {
				//广播的数据有时候会没有移动或者没有方向，此时地址会出错
				roomFrames[i*FramesPerBag+j] = Frame{GetDeltaDirection(send.ClientInfo[i].Msg[j].Move),
					send.ClientInfo[i].Msg[j].DeltaTime,
					send.ClientInfo[i].Msg[j].Skillid}
			}
		}
		str := "room.roomframes"
		for i := int32(0); i < RoomPeople*FramesPerBag; i++ {
			c.Update(bson.M{"roomid": roomid}, bson.M{"$push": bson.M{str: roomFrames[i]}})
		}
	}
}

//todo 可以将存储的房间对战数据取出用于回放功能
func GetFrame() {

}

//广播的数据有时候会没有移动或者没有方向，此时是nil
func GetDeltaDirection(deltaDirection *DTO.DeltaDirection) DeltaDirection {
	if deltaDirection == nil {
		return DeltaDirection{0, 0}
	} else {
		return DeltaDirection{deltaDirection.X, deltaDirection.Y}
	}
}
