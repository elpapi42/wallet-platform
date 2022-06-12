package gin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	server *http.Server
}

func NewGinServer(port int) *GinServer {
	engine := gin.Default()

	RegisterRoutes(engine)

	server := &GinServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: engine,
		},
	}

	return server
}

func (g *GinServer) Start() {
	go g.Listen()
}

func (g *GinServer) Listen() {
	err := g.server.ListenAndServe()

	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Println("gin server closed:", err)
	}
}

func (g *GinServer) Stop() {
	log.Println("stopping gin server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := g.server.Shutdown(ctx)
	if err != nil {
		log.Println("gin server forced to shutdown:", err)
	}
	log.Println("gin server stopped")
}
