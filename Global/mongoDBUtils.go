package Global

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

//数据库结构
//User层	userid(int32	start 100000)	uuid(string)
//Room层  Room{
//			Roomid{
//				playerid	seatid		role
//				bagid{
// 					seatid{
// 						frameid{
// 							move
// 							skill
// 							}
// 					}
// 				}
// }
// }
//
//{"_id":"5d5281677ebff410283c89f0","RoomID":1,"Room":{"Players":[{"playerId":1,"seatId":1,"RoleID":1}],"Frames":[{"Move":{"X":1,"Y":1},"SkillID":1,"Skill":{"X":1,"Y":1}}]}}
//
type MongoDBRoom struct {
}

func ConnecToDB() *mgo.Collection {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("Mirror").C("Room")
	fmt.Println("Session ok")
	//c.Insert({"_id":"5d5281677ebff410283c89f0","RoomID":1,"Room":{"Players":[{"playerId":1,"seatId":1,"RoleID":1}],"Frames":[{"Move":{"X":1,"Y":1},"SkillID":1,"Skill":{"X":1,"Y":1}}]}})
	//c.Insert({"5d5281677ebff410283c89f0",1,{ [ {1,1,1} ], [{ {1,1},1, {1,1} }]}})
	return c
}
