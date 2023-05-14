package delivery

import (
	"context"
	"fmt"
	"github.com/morris-zheng/go-slim/internal/delivery/user"
	"github.com/morris-zheng/go-slim/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
	engine *gin.Engine
}

func NewHttpServer(svc *domain.ServiceContext) *HttpServer {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	return &HttpServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", svc.Config.Port),
			Handler: e,
		},
		engine: e,
	}
}

func (s *HttpServer) Run(ctx context.Context, svc *domain.ServiceContext) {
	go func() {
		<-ctx.Done()
		svc.Logger.Info(ctx, "Shutdown Server ...")
		wait, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if err := s.server.Shutdown(wait); err != nil {
			svc.Logger.Fatal(ctx, fmt.Sprintf("Server Shutdown: %v", err))
		}
		<-wait.Done()
	}()

	svc.Logger.Info(ctx, fmt.Sprintf("server listening at: %d", svc.Config.Port))
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		svc.Logger.Fatal(ctx, fmt.Sprintf("listen err: %v", err))
	}
}

func (s *HttpServer) Register(svc *domain.ServiceContext) {
	// inject handler
	userHandler := user.NewHandler(svc)

	// route info
	userGroup := s.engine.Group("/user")
	{
		userGroup.GET("", userHandler.Query)
		userGroup.GET("/:id", userHandler.Get)
		userGroup.POST("", userHandler.Create)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.DELETE("/:id", userHandler.Delete)
	}
}
