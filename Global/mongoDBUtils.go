package Global

import (
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"../proto/dto"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

const DataBase string = "Mirror"
const RoomDB string = "Room"
const UserDB string = "User"

//user数据库
type User struct {
	playerid int32
	uuid     string
}

func GetUser(uuid string) {
	//c := ConnecToDB(UserDB)
	//user := User{}
	//err := c.Find(bson.M{"uuid": uuid}).One(&user)
}

//仅供GetUser调用,故为包内可见
func insertUser() {

}

//room数据库
type MongoDBRoomMng struct {
	_id    string
	RoomID int32
	Room   MongoDBRoom
}

type MongoDBRoom struct {
	RoomPlayers []Player
	RoomFrames  []Frame
}
type Player struct {
	playerid int32
	name     string
	roleid   int32
	seat     int32
}

type Frame struct {
	FrameMove    DeltaDirection
	FrameSkill   DeltaDirection
	FrameSkillID int32
}

type DeltaDirection struct {
	x         float32
	y         float32
	deltatime float32
}

func ConnecToDB(DB string) *mgo.Collection {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DataBase).C(DB)
	fmt.Println("Session ok")
	return c
}

//新开一个房间就新建一个Room数据
func AddRoom(roomid int32, match *DTO.MatchSuccessDTO) {
	c := ConnecToDB(RoomDB)

	//player信息房间匹配好就传进来建立表
	players := make([]Player, RoomPeople)
	//frame := make([]Frame, RoomPeople*FramesPerBag)

	for i := int32(0); i < RoomPeople; i++ {
		players[i] = Player{match.Players[i].Playerid, match.Players[i].Name, match.Players[i].Roleid, match.Players[i].Seat}
	}
	room := MongoDBRoom{players, []Frame{}}
	roomMng := MongoDBRoomMng{"1", roomid, room}
	c.Insert(roomMng)
	fmt.Println("add room ok")
}

//每次插入RoomPeople*FramesPerBag帧
func AddFrame(roomid int32, send *DTO.ServerMoveDTO) {
	c := ConnecToDB(RoomDB)
	roomFrames := make([]Frame, RoomPeople*FramesPerBag)
	fmt.Println("start add frame ")

	for i := int32(0); i < RoomPeople; i++ {
		for j := int32(0); j < FramesPerBag; j++ {
			//广播的数据有时候会没有移动或者没有方向，此时地址会出错
			roomFrames[i*FramesPerBag+j] = Frame{GetDeltaDirection(send.ClientInfo[i].Msg[j].Move),
				GetDeltaDirection(send.ClientInfo[i].Msg[j].SkillDir),
				send.ClientInfo[i].Msg[j].Skillid}
		}
	}
	str := "room.roomframes"
	for i := int32(0); i < RoomPeople*FramesPerBag; i++ {
		c.Update(bson.M{"roomid": roomid}, bson.M{"$push": bson.M{str: roomFrames[i]}})
	}
	fmt.Println("start add frame ok")
}

//todo 可以将存储的房间对战数据取出用于回放功能
func GetFrame() {

}

//广播的数据有时候会没有移动或者没有方向，此时是nil
func GetDeltaDirection(deltaDirection *DTO.DeltaDirection) DeltaDirection {
	if deltaDirection == nil {
		return DeltaDirection{0, 0, 0}
	} else {
		return DeltaDirection{deltaDirection.X, deltaDirection.Y, deltaDirection.DeltaTime}
	}
}
