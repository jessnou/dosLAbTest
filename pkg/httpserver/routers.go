package httpserver

import (
	"dosLAbTest/pkg/postgres/query"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (server *Server) getStatistics(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		fmt.Errorf("cannot convert string to int", err)
	}
	response, err := query.GetPost(server.store, postId)
	if err != nil {
		fmt.Errorf("cannot get post", err)
	}

	ctx.JSON(http.StatusOK, response)
}
