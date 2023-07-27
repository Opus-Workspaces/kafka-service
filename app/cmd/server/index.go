package server

import (
	"context"
	"fmt"
	"kafka-service/app/config"
	dbCfg "kafka-service/app/config/database"
	serverCfg "kafka-service/app/config/server"

	"kafka-service/app/models"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo

	ServerConfig serverCfg.ServerType
	DBConfig     dbCfg.DatabaseType

	// DBRedis *redis.Client
}

func InitServer(cfg config.Config) *Server {

	serverCfg := cfg.Server
	// dbCfg := cfg.DB

	// mongoConn := mongo.ConnectMongo(dbCfg)
	// redisConn := redis.NewRedisClient(dbCfg)

	return &Server{
		Echo:         echo.New(),
		ServerConfig: serverCfg,
		// DBConfig:     dbCfg,
		// DBMongo:      mongoConn,
		// DBRedis:      redisConn,
	}
}

func Run(s *Server) {
	e := s.Echo
	rateLimiter := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      10,
				Burst:     30,
				ExpiresIn: 1 * time.Minute,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, models.ResponseError{
				StatusCode: http.StatusForbidden,
				Type:       "server.go.rate_limiter.error_handler",
				Message:    err.Error(),
			})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, models.ResponseError{
				StatusCode: http.StatusTooManyRequests,
				Type:       "server.go.rate_limiter.deny_handler",
				Message:    err.Error(),
			})
		},
	}

	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders},
	}))
	e.Use(middleware.RateLimiterWithConfig(rateLimiter))
	s.Echo.GET("/", func(e echo.Context) error {
		time.Sleep(5 * time.Second) // Simulate long running task
		return e.JSON(http.StatusOK, "Hello World!")
	})

	serverConfig := s.ServerConfig

	go func() {
		if err := e.Start(serverConfig.Port); err != nil && err != http.ErrServerClosed {
			fmt.Println("error starting server: ", err)
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}
}
