package mgo

import (
	"fmt"
	"log"
	"os"

	"github.com/globalsign/mgo"
)

// 定义mongo session
var globalS *mgo.Session

//初始化连接MongoDB
func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("Addrs")},
		Source:   os.Getenv("source"),
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error", err)
	}
	fmt.Println("MongoDB Connect")
	globalS = session
}

//连接MongoDB 返回一个session会话和一个集合c
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

//插入
func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

//查找某一个函数
func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

//是否存在
func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

//查找所有
func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

//更新
func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

//删除
func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
