package domain

import (
	"fmt"
	"github.com/morris-zheng/go-slim/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config  *conf.Config
	DB      *gorm.DB
	DBSlave *gorm.DB
}

func NewServiceContext(c *conf.Config) *ServiceContext {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm init err", err)
	}
	if c.Debug {
		db = db.Debug()
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
