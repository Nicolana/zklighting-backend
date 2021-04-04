package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lexkong/log"

	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	log.Info(config)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	// setupDB(db)
	return db
}

// func setupDB(db *gorm.DB) {
// 	db.LogMode(viper.GetBool("gormlog"))
// 	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
// 	db.DB().SetMaxIdleConns(0)
// }

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// 为Database结构的类型添加方法
func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}
