package markets

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"github.com/h-varmazyar/gopack/mux"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	marketsService chipmunkApi.MarketServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			marketsService: chipmunkApi.NewMarketServiceClient(chipmunkConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	markets := router.PathPrefix("/markets").Subrouter()

	markets.HandleFunc("/create", c.create).Methods(http.MethodPost)
	markets.HandleFunc("/list", c.list).Methods(http.MethodGet)
	markets.HandleFunc("/{market-id}", c.get).Methods(http.MethodGet)
	markets.HandleFunc("/{market-id}", c.update).Methods(http.MethodPut)
	markets.HandleFunc("/{market-id}", c.delete).Methods(http.MethodDelete)
}

func (c Controller) create(res http.ResponseWriter, req *http.Request) {
	market := new(chipmunkApi.CreateMarketReq)
	if err := httpext.BindModel(req, market); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.marketsService.Create(req.Context(), market); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}

func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	list := new(chipmunkApi.MarketListRequest)
	brokerageIDs := mux.QueryParam(req, "brokerage-id")
	if len(brokerageIDs) != 0 {
		list.BrokerageID = brokerageIDs[0]
	}
	if markets, err := c.marketsService.List(req.Context(), list); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, markets.Elements)
	}
}

func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getRequest := new(chipmunkApi.ReturnMarketRequest)
	getRequest.ID = mux.PathParam(req, "market-id")

	if market, err := c.marketsService.Return(req.Context(), getRequest); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, market)
	}
}

func (c Controller) update(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusNotFound)
}

func (c Controller) delete(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusNotFound)
}
