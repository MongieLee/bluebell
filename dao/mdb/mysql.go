package mdb

import (
	"bluebell/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	fmt.Println(cfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	_db, err := db.DB()
	if err == nil {
		_db.SetMaxOpenConns(cfg.MaxOpenConns)
		_db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	return
}

func Close() {
	_db, err := db.DB()
	if err == nil {
		_db.Close()
	}
}
