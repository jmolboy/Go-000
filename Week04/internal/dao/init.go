package dao


import (
	"fmt"
	"gorm.io/driver/mysql" // 引入mysql驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


// MysqlInvoker MySQL调用器。
var MyDB *gorm.DB

// NewMysqlInvoker 返回新的mysql invoker
func NewMysql() (invoker *gorm.DB) {
	fmt.Println("=============begin mysql============")
	invoker = newMysql()
	fmt.Println("=============end mysql============")

	return
}

// Init 初始化model。
func Init() {
	MyDB = NewMysql()
}

func newMysql() (db *gorm.DB) {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // 禁用彩色打印
		},
	)

	var dsn = ""

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger})

	if err != nil {
		panic("failed to connect mysql:" + dsn + ", error: " + err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)

	if err != nil {
		panic("failed to connect mysql:" + dsn)
	}
	return
}
