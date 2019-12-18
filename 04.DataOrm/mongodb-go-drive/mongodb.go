// mongodb官方推出的go语言驱动器，相比较mgo来说，使用较为复杂一点
package mongodb_go_drive

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func InitMongoDB() {
	// 配置client
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_HOST"))
	// 连接mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// 检查是否连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	MongoDB = client
}

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoDB.Database("demo").Collection("user")

	// 新增单个数据
	if result, err := db.InsertOne(ctx, User{Name: "UserName", Phone: "1234567890"}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// 批量新增数据
	{
		users := []interface{}{
			User{Name: "UserName_0", Phone: "123"},
			User{Name: "UserName_1", Phone: "456"},
			User{Name: "UserName_2", Phone: "789"},
		}
		if result, err := db.InsertMany(ctx, users); err == nil {
			log.Println(result)
		} else {
			log.Fatal(err)
		}
	}

	// 查询单个
	{
		result := db.FindOne(ctx, bson.M{"phone": "1234567890"})
		var user User
		if err := result.Decode(&user); err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}

	// 查询
	if cur, err := db.Find(ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}}); err == nil {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var user User
			if err := cur.Decode(&user); err != nil {
				log.Fatal(err)
			}
			log.Println(user)
		}
	} else {
		log.Fatal(err)
	}

	// 修改单个
	if result, err := db.UpdateOne(
		ctx, bson.M{"phone": "123"},
		bson.M{"$set": bson.M{"dbname": "UserName_changed"}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// 更新多个
	if result, err := db.UpdateMany(
		ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}},
		bson.M{"$set": bson.M{"dbname": "UserName_changed"}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// 替换
	{
		user := User{Name: "UserName_2_replaced", Phone: "789"}
		if result, err := db.ReplaceOne(ctx, bson.M{"phone": "789"}, user); err == nil {
			log.Println(result)
		} else {
			log.Fatal(err)
		}
	}

	// 删除一个
	if result, err := db.DeleteOne(ctx, bson.M{"phone": "123"}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// 批量删除
	if result, err := db.DeleteMany(ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
	//使用完成之后需要关闭连接
	err := MongoDB.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID    uint
	Name  string `bson:"name",json:"name"`
	Phone string
}
