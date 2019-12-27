package xorm

import (
	"fmt"
	"log"
)

type UserTbl struct {
	Id       int
	Username string
	Sex      int
	Info     string
}

func selectAll() {
	var user []UserTbl
	engine.SQL("select * from user_tbl").Find(&user)
	fmt.Println(user)
}

//条件查询
func selectUser(name string) {
	var user []UserTbl
	engine.Where("user_tbl.username=?", name).Find(&user)
	fmt.Println(user)
}

//可以用Get查询单个元素
func selectOne(id int) {
	var user UserTbl
	engine.Id(id).Get(&user)
	//engine.Alias("u").Where("u.id=?",id).Get(&user)
	fmt.Println(user)
}

//添加
func InsertUser(user *UserTbl) bool {
	rows, err := engine.Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//删除(根据名称删除)
func DeleteUser(name string) bool {
	user := UserTbl{
		Username: name,
	}
	rows, err := engine.Delete(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//利用sql删除
func DeleteUserBySQL(name string) bool {
	result, err := engine.Exec("delete from user_tbl where username=?", name)
	if err != nil {
		log.Println(err)
		return false
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true
	}
	return false
}

//更新
func UpdateUser(user *UserTbl) bool {
	//Update(bean interface{}, condiBeans ...interface{}) bean是需要更新的bean,condiBeans是条件
	rows, err := engine.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		log.Println(err)
		return false
	}
	if rows > 0 {
		return true
	}
	return false
}

//利用session进行增删改
//用session的好处就是可以事务处理
func SessionUserTest(user *UserTbl) {
	session := engine.NewSession()
	session.Begin()
	_, err := session.Insert(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	user.Username = "windows"
	_, err = session.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	_, err = session.Delete(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	err = session.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
