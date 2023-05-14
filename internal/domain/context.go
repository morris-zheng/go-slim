package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/morris-zheng/go-slim/internal/conf"

	"github.com/morris-zheng/go-slim-core/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  *conf.Config
	Logger  logger.Logger
	DB      *gorm.DB
	DBSlave *gorm.DB
}

var svc *ServiceContext

func NewServiceContext(c *conf.Config) *ServiceContext {
	if svc == nil {
		// logger
		l, err := logger.NewLogger(logger.Level(c.Logger.Level))
		if err != nil {
			log.Fatal("logger init err: ", err)
		}

		// db
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			l.Fatal(context.Background(), fmt.Sprintf("DB init err: %v", err))
		}

		if c.Debug {
			db = db.Debug()
		}

		svc = &ServiceContext{
			Config: c,
			Logger: l,
			DB:     db,
		}
	}

	return svc
}
