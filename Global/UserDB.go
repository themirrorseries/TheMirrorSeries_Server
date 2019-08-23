package Global

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ConnecToUser() *mgo.Collection {
	session, err := mgo.Dial(MongoURL)
	if err != nil {
		ErrorLog.Log.Errorln("User数据库连接出错！")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DataBase).C(UserDB)
	return c
}

func GetUser(c *mgo.Collection, uuid string) int32 {
	user := MongoDBUser{}
	err := c.Find(bson.M{"uuid": uuid}).One(&user)
	if err != nil {
		NextUserIDMu.Lock()
		ret := NextUserID
		c.Insert(&MongoDBUser{NextUserID, uuid})
		NextUserID++
		NextUserIDMu.Unlock()
		DetailedLog.Log.Info("该设备第一次连接，将为设备分配唯一用户ID---...", ret)
		return ret
	} else {
		DetailedLog.Log.Info("已查询到该设备用户ID---...", user.Playerid)
		return user.Playerid
	}
}

//服务端启动的时候读取数据库最后一个UserID
func GetLastUserID(c *mgo.Collection) int32 {
	count, _ := c.Count()
	if count == 0 {
		return 100000
	} else {
		fmt.Println(count)
		var a MongoDBUser
		c.Find(nil).Sort("-playerid").One(&a)
		return a.Playerid + 1
	}
}
