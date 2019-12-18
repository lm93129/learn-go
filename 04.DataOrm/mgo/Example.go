package mgo

import "github.com/globalsign/mgo/bson"

const (
	db         = "demo" // 设置数据库名
	collection = "UserModel"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Phone    string        `json:"phone" bson:"phone"`
	Password string        `json:"-" bson:"passoword"`
	EmpName  string        `json:"empname" bson:"empname"`
	Role     string        `json:"role" bson:"role"`
}

//新增用户
func (e *User) AddEmployee(emp User) error {
	return Insert(db, collection, emp)
}

//查找用户
func (e *User) FindUserByPhone(phone string) (User, error) {
	var result User
	err := FindOne(db, collection, bson.M{
		"phone": phone,
	}, nil, &result)
	return result, err
}
