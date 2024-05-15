package http_hdl

import (
	"fmt"

	"github.com/PaoloEG/terrasense/internal/core/services"
	"github.com/PaoloEG/terrasense/internal/infrastructure/handlers/http/handlers"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	server    *gin.Engine
	msService *services.MeasurementsService
}

func New(msSrv *services.MeasurementsService, isProd bool) *HttpHandler {
	if isProd {
		gin.SetMode("release")
	}
	return &HttpHandler{
		server:    gin.Default(),
		msService: msSrv,
	}
}

func (h *HttpHandler) Start(port string) {
	// Add routes
	h.server.GET("/telemetries/:chipID/latest", handlers.GetLatestHandler(h.msService))
	h.server.GET("/telemetries/:chipID", handlers.GetRangeHandler(h.msService))
	//start server
	h.server.Run(fmt.Sprintf(":%s", port))
}
