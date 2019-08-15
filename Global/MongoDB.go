package Global

import ()

const DataBase string = "Mirror"
const RoomDB string = "Room"
const UserDB string = "User"

//user数据库
type MongoDBUser struct {
	Playerid int32
	Uuid     string
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
	Playerid int32
	Name     string
	Roleid   int32
	Seat     int32
}

type Frame struct {
	FrameMove    DeltaDirection
	FrameSkill   DeltaDirection
	FrameSkillID int32
}

type DeltaDirection struct {
	X         float32
	Y         float32
	Deltatime float32
}
