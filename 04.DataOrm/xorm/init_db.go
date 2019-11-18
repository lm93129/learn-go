package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/core"
)

var engine *xorm.Engine

// 单数据库初始化
func Database() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123@127.0.0.1:3306/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err := engine.Ping(); err != nil {
		panic(err)
	}
	// 初始化sqlite3数据库
	// engine, err = xorm.NewEngine("sqlite3", "./test.db")
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})
	// 将执行的sql语句打印出来
	engine.ShowSQL(true)

	// 同步数据库
	migration()

}

// 初始化数据库集群
var eg *xorm.EngineGroup

// 默认的负载策略为轮询负载策略
func DataBaseGroup() {
	conns := []string{
		"postgres://postgres:root@localhost:5432/test?sslmode=disable;",  // 第一个默认是master
		"postgres://postgres:root@localhost:5432/test1?sslmode=disable;", // 第二个开始都是slave
		"postgres://postgres:root@localhost:5432/test2?sslmode=disable",
	}

	var err error
	eg, err = xorm.NewEngineGroup("postgres", conns) //默认的轮询负载策略
	// eg, err = xorm.NewEngineGroup("postgres", conns, xorm.RandomPolicy()) //随机访问负载策略
	// eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRandomPolicy([]int{2, 3})) //随机权重访问负载策略
	// eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3})) //权重轮询访问负载策略
	// eg, err = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy())                   //最小连接数访问负载策略

	if err != nil {
		panic(err)
	}
	eg.ShowSQL(true)

}

// 另外一种初始化数据库的方式
func DataBaseGroup2() {
	var err error
	master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	if err != nil {
		return
	}

	slave1, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test1?sslmode=disable")
	if err != nil {
		return
	}

	slave2, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test2?sslmode=disable")
	if err != nil {
		return
	}

	slaves := []*xorm.Engine{slave1, slave2}
	eg, err = xorm.NewEngineGroup(master, slaves)
	if err != nil {
		panic(err)
	}
	eg.ShowSQL(true)
}

// 创建完成EngineGroup之后，并没有立即连接数据库，
// 此时可以通过eg.Ping()来进行数据库的连接测试是否可以连接到数据库，该方法会依次调用引擎组中每个Engine的Ping方法。
// 另外对于某些数据库有连接超时设置的，可以通过起一个定期Ping的Go程来保持连接鲜活。
// EngineGroup可以通过eg.Close()来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。
