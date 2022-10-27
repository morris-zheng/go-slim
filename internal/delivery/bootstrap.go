package delivery

import (
	"github.com/morris-zheng/go-slim/internal/delivery/user"
	"github.com/morris-zheng/go-slim/internal/domain"

	"github.com/gin-gonic/gin"
)

func Register(svc *domain.ServiceContext, r *gin.Engine) {
	// inject handler
	userHandler := user.NewHandler(svc)

	// route info
	userGroup := r.Group("/user")
	{
		userGroup.GET("", userHandler.Query)
		userGroup.GET("/:id", userHandler.Get)
		userGroup.POST("", userHandler.Create)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.DELETE("/:id", userHandler.Delete)
	}
}
