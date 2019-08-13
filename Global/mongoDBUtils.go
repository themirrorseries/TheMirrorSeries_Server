package Global

import (
	"gopkg.in/mgo.v2"
)

//数据库结构
//User层	userid(int32	start 100000)	uuid(string)
//Room层  Room{
//			Roomid{
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
//
//
func ConnecToDB() *mgo.Collection {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("medex").C("student")
	return c
}
