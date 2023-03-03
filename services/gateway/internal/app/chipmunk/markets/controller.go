package markets

import (
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
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
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	markets := router.PathPrefix("/markets").Subrouter()

	markets.HandleFunc("/create", c.create).Methods(http.MethodPost)
	markets.HandleFunc("/list", c.list).Methods(http.MethodGet)
	markets.HandleFunc("/UpdateDetails", c.updateDetails).Methods(http.MethodPost)
	markets.HandleFunc("/Update", c.update).Methods(http.MethodPost)
	//markets.HandleFunc("/StartWorker", c.startWorker).Methods(http.MethodPost)
	//markets.HandleFunc("/StopWorker", c.stopWorker).Methods(http.MethodPost)
	markets.HandleFunc("/{market-id}", c.get).Methods(http.MethodGet)
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
	list := new(chipmunkApi.MarketListReq)
	platforms := mux.QueryParam(req, "Platform")
	if len(platforms) != 0 {
		list.Platform = api.Platform(api.Platform_value[platforms[0]])
	}
	if markets, err := c.marketsService.List(req.Context(), list); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, markets.Elements)
	}
}

func (c Controller) updateDetails(res http.ResponseWriter, req *http.Request) {
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

//func (c Controller) startWorker(res http.ResponseWriter, req *http.Request) {
//	worker := new(chipmunkApi.WorkerStartReq)
//	if err := httpext.BindModel(req, worker); err != nil {
//		httpext.SendError(res, req, err)
//		return
//	}
//	if _, err := c.marketsService.StartWorker(req.Context(), worker); err != nil {
//		httpext.SendError(res, req, err)
//	} else {
//		httpext.SendCode(res, req, http.StatusOK)
//	}
//}
//
//func (c Controller) stopWorker(res http.ResponseWriter, req *http.Request) {
//	worker := new(chipmunkApi.WorkerStopReq)
//	if err := httpext.BindModel(req, worker); err != nil {
//		httpext.SendError(res, req, err)
//		return
//	}
//	if _, err := c.marketsService.StopWorker(req.Context(), worker); err != nil {
//		httpext.SendError(res, req, err)
//	} else {
//		httpext.SendCode(res, req, http.StatusOK)
//	}
//}

func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getRequest := new(chipmunkApi.MarketReturnReq)
	getRequest.ID = mux.PathParam(req, "market-id")

	if market, err := c.marketsService.Return(req.Context(), getRequest); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, market)
	}
}

func (c Controller) update(res http.ResponseWriter, req *http.Request) {
	update := new(chipmunkApi.MarketUpdateReq)
	if err := httpext.BindModel(req, update); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if markets, err := c.marketsService.Update(req.Context(), update); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, markets.Elements)
	}
}

func (c Controller) delete(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusNotFound)
}
