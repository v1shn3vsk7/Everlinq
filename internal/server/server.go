package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/v1shn3vsk7/Everlinq/internal/config"
	"github.com/v1shn3vsk7/Everlinq/internal/data/db/mongodb"
	"github.com/v1shn3vsk7/Everlinq/internal/server/api/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	cfg        *config.Config
	engine     *gin.Engine
	db         *mongo.Database
}

func NewServer(config *config.Config) *Server {
	return &Server{
		cfg:    config,
		engine: gin.New(),
		db:     mongodb.NewClient(*config.Context, config.DbUrl),
	}
}

func (s *Server) Run(ctx context.Context) error {
	gin.SetMode(s.cfg.LogLevel)

	s.engine.Use(gin.Recovery())

	s.setupRoutes(ctx)

	s.engine.Static("/css", "./templates/css")

	s.engine.LoadHTMLGlob("templates/*.html")

	return s.engine.Run(s.cfg.BindAddr)
}

func (s *Server) setupRoutes(ctx context.Context) {
	s.engine.GET("/", handlers.Index())
	s.engine.GET("/login", handlers.Login())
}
