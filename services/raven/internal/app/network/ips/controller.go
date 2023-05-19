package ips

import (
	"fmt"
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	ipsService networkApi.IPServiceClient
	logger     *log.Logger
}

func ControllerInstance(logger *log.Logger, networkAddress string) *Controller {
	if controller == nil {
		networkConn := grpcext.NewConnection(networkAddress)
		controller = &Controller{
			ipsService: networkApi.NewIPServiceClient(networkConn),
			logger:     logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	ips := router.PathPrefix("/ips").Subrouter()

	ips.HandleFunc("/create", c.create).Methods(http.MethodPost)
	ips.HandleFunc("/list", c.list).Methods(http.MethodGet)
	ips.HandleFunc("/{ip_id}", c.returnByID).Methods(http.MethodGet)
}

// ipCreate godoc
//	@Summary		Create new IP
//	@Description	Create new IP
//	@Accept			json
//	@Produce		json
//	@Param			IP	body		proto.IPCreateReq	true	"New IP"
//	@Success		201	{object}	proto.IP
//	@Failure		400	{object}	errors.Error
//	@Failure		404	{object}	errors.Error
//	@Failure		500	{object}	errors.Error
//	@Router			/network/ips/create [post]
func (c Controller) create(res http.ResponseWriter, req *http.Request) {
	fmt.Println("ip creation")
	ipCreateReq := new(networkApi.IPCreateReq)
	if err := httpext.BindModel(req, ipCreateReq); err != nil {
		httpext.SendError(res, req, err)
		return
	}

	fmt.Println("create:", ipCreateReq)
	if ip, err := c.ipsService.Create(req.Context(), ipCreateReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusCreated, ip)
	}
}

// ipList godoc
//	@Summary		return IP list
//	@Description	return IP list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	proto.IPs
//	@Failure		400	{object}	errors.Error
//	@Failure		404	{object}	errors.Error
//	@Failure		500	{object}	errors.Error
//	@Router			/network/ips/list [get]
func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	fmt.Println("ip list req")
	if ips, err := c.ipsService.List(req.Context(), new(networkApi.IPListReq)); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, ips)
	}
}

// ipReturn godoc
//	@Summary		return IP with ID
//	@Description	return IP with ID
//	@Accept			json
//	@Produce		json
//	@Param			ip_id	path		string	true	"IP ID"
//	@Success		200		{object}	proto.IP
//	@Failure		400		{object}	errors.Error
//	@Failure		404		{object}	errors.Error
//	@Failure		500		{object}	errors.Error
//	@Router			/network/ips/{ip_id} [get]
func (c Controller) returnByID(res http.ResponseWriter, req *http.Request) {
	ipReq := new(networkApi.IPReturnReq)

	ipReq.ID = mux.PathParam(req, "ip_id")

	fmt.Println("ip return:", ipReq.ID)

	if ip, err := c.ipsService.Return(req.Context(), ipReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, ip)
	}
}
