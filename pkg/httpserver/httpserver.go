package httpserver

import (
	"dosLAbTest/config"
	"dosLAbTest/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  postgres.Postgres
	router *gin.Engine
}

func NewServer(config config.Config, store postgres.Postgres) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/post/:postId/comments", server.getStatistics)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
