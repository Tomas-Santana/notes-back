package api

import (
	"notes-back/database"

	"notes-back/api/auth"
	"notes-back/api/resource"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	listenAddr string
	database   database.Database
	engine     *gin.Engine
	validator  *validator.Validate
}

func NewServer(listenAddr string, database database.Database, validator *validator.Validate) *Server {
	return &Server{
		listenAddr: listenAddr,
		database:   database,
		engine:     gin.Default(),
		validator:  validator,
	}
}
func (s *Server) NewGroup(path string) *gin.RouterGroup {
	return s.engine.Group(path)
}

func (s *Server) Start() error {
	s.CreateRoutes()
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return s.engine.Run(s.listenAddr)

}

func (s *Server) CreateRoutes() {
	authGroup := s.NewGroup("/auth")
	authRouter := auth.NewAuthRouter(s.database, authGroup, s.validator)
	authRouter.RegisterRoutes()

	resourceGroup := s.NewGroup("/resource")
	resourceRouter := resource.NewRouter(s.database, resourceGroup, s.validator)
	resourceRouter.RegisterRoutes()
}
