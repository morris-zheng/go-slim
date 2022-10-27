package user

import (
	"github.com/morris-zheng/go-slim/internal/domain"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(ctx *domain.ServiceContext) *Repo {
	return &Repo{
		DB: ctx.DB,
	}
}
