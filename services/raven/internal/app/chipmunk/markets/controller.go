package markets

import (
	"fmt"
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	marketsService chipmunkApi.MarketServiceClient
	logger         *log.Logger
}

func ControllerInstance(logger *log.Logger, chipmunkAddress string) *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(chipmunkAddress)
		controller = &Controller{
			marketsService: chipmunkApi.NewMarketServiceClient(chipmunkConn),
			logger:         logger,
		}
	}
	fmt.Println("controller created")
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	markets := router.PathPrefix("/markets").Subrouter()

	markets.HandleFunc("/create", c.create).Methods(http.MethodPost)
	markets.HandleFunc("/list", c.list).Methods(http.MethodGet)
	markets.HandleFunc("/{market_id}", c.update).Methods(http.MethodPut)
	markets.HandleFunc("/{market-id}", c.get).Methods(http.MethodGet)
	markets.HandleFunc("/Update-remotely", c.updateDetails).Methods(http.MethodPost)

	fmt.Println("route registered")

}

// marketCreate godoc
//
//	@Summary		Create new market manually
//	@Description	Create new market manually
//	@Accept			json
//	@Produce		json
//	@Param			market	body	MarketReq	true	"New Market"
//	@Success		201
//	@Failure		400	{object}	errors.Error
//	@Failure		404	{object}	errors.Error
//	@Failure		500	{object}	errors.Error
//	@Router			/chipmunk/markets/create [post]
func (c Controller) create(res http.ResponseWriter, req *http.Request) {
	market := new(chipmunkApi.MarketCreateReq)
	if err := httpext.BindModel(req, market); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	c.logger.Infof("filled market is: %v", market)
	if _, err := c.marketsService.Create(req.Context(), market); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}

// marketList godoc
//
//	@Summary		get market list
//	@Description	get market list based on platform
//	@Accept			json
//	@Produce		json
//	@Param			platform	query		string	true	"Platform name"	Enums:(Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance)
//	@Success		200			{object}	proto.Markets
//	@Failure		400			{object}	errors.Error
//	@Failure		404			{object}	errors.Error
//	@Failure		500			{object}	errors.Error
//	@Router			/chipmunk/markets/list [get]
func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	fmt.Println("market list")
	list := new(chipmunkApi.MarketListReq)
	platforms := mux.QueryParam(req, "platform")
	if len(platforms) == 0 {
		httpext.SendError(res, req, errors.New(req.Context(), codes.InvalidArgument).AddDetails("platform needed"))
		return
	}
	if platforms[0] != "" {
		list.Platform = api.Platform(api.Platform_value[platforms[0]])
	}

	if markets, err := c.marketsService.List(req.Context(), list); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, markets.Elements)
	}
}

// marketsUpdateRemotely godoc
//
//	@Summary		update markets remotely
//	@Description	update markets remotely
//	@Accept			json
//	@Produce		json
//	@Param			market	body		Platform	true	"update Markets"
//	@Success		200		{object}	proto.Markets
//	@Failure		400		{object}	errors.Error
//	@Failure		404		{object}	errors.Error
//	@Failure		500		{object}	errors.Error
//	@Router			/chipmunk/markets/Update-remotely [post]
func (c Controller) updateDetails(res http.ResponseWriter, req *http.Request) {
	fmt.Println("update details")
	update := new(chipmunkApi.MarketUpdateFromPlatformReq)
	if err := httpext.BindModel(req, update); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if markets, err := c.marketsService.UpdateFromPlatform(req.Context(), update); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, markets.Elements)
	}
}

// marketReturn godoc
//
//	@Summary		return market
//	@Description	return market based on id
//	@Accept			json
//	@Produce		json
//	@Param			market_id	path		string	true	"market id"
//	@Success		200			{object}	proto.Market
//	@Failure		400			{object}	errors.Error
//	@Failure		404			{object}	errors.Error
//	@Failure		500			{object}	errors.Error
//	@Router			/chipmunk/markets/{market_id} [get]
func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getRequest := new(chipmunkApi.MarketReturnReq)
	getRequest.ID = mux.PathParam(req, "market-id")

	if market, err := c.marketsService.Return(req.Context(), getRequest); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, market)
	}
}

// marketUpdate godoc
//
//	@Summary		update new market manually
//	@Description	update new market manually
//	@Accept			json
//	@Produce		json
//	@Param			market_id	path	string		true	"market id"
//	@Param			market		body	MarketReq	true	"update Market"
//	@Success		201
//	@Failure		400	{object}	errors.Error
//	@Failure		404	{object}	errors.Error
//	@Failure		500	{object}	errors.Error
//	@Router			/chipmunk/markets/{market_id} [put]
func (c Controller) update(res http.ResponseWriter, req *http.Request) {
	update := new(chipmunkApi.MarketUpdateReq)
	if err := httpext.BindModel(req, update); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	update.ID = mux.PathParam(req, "market_id")
	if market, err := c.marketsService.Update(req.Context(), update); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, market)
	}
}
