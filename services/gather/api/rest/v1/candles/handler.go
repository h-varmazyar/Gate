package candles

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/services/gather/internal/services/candles"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	logger        *log.Logger
	candleService candles.Service
}

func New(logger *log.Logger, candlesService candles.Service) Handler {
	return Handler{
		logger:        logger,
		candleService: candlesService,
	}
}

func (h *Handler) List(c *gin.Context) {
	marketIDParam := c.Param("market-id")
	resolutionIDParam := c.Param("resolution-id")

	marketID, err := strconv.ParseUint(marketIDParam, 10, 64)
	if err != nil {
		httpext.SendGinError(c, err)
		return
	}

	resolutionID, err := strconv.ParseUint(resolutionIDParam, 10, 64)
	if err != nil {
		httpext.SendGinError(c, err)
		return
	}

	resp, err := h.candleService.List(c, uint(marketID), uint(resolutionID))
	if err != nil {
		httpext.SendGinError(c, err)
		return
	}

	httpext.SendGinModel(c, http.StatusOK, resp)
}
