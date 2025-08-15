package main

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/Gate/pkg/db"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/logger"
	v1 "github.com/h-varmazyar/Gate/services/gather/api/rest/v1"
	"github.com/h-varmazyar/Gate/services/indicators/configs"
	"github.com/h-varmazyar/Gate/services/indicators/internal/router"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"os"
	"time"
)

type dependencies struct {
	ctx          context.Context
	StopSignal   chan os.Signal
	Log          *log.Logger
	Cfg          *configs.Configs
	DB           *gorm.DB
	Gin          *gin.Engine
	Repositories struct {
	}
	Services struct {
	}
	Controllers struct {
	}
	Routers struct {
		Router      router.Router
		V1          v1.Router
		Middlewares struct {
		}
	}
	Workers struct {
	}
}

func InjectDependencies() (dep dependencies, err error) {
	if err = generalDependencies(&dep); err != nil {
		return
	}

	if err = infraDependencies(&dep); err != nil {
		return
	}

	if err = repositoryDependencies(&dep); err != nil {
		return
	}

	if err = serviceDependencies(&dep); err != nil {
		return
	}

	if err = controllerDependencies(&dep); err != nil {
		return
	}

	if err = routersDependencies(&dep); err != nil {
		return
	}

	if err = workerDependencies(&dep); err != nil {
		return
	}

	dep.Log.Info("Injecting dependencies")

	return
}

var generalDependencies = func(dep *dependencies) (err error) {
	dep.ctx = context.Background()
	dep.StopSignal = make(chan os.Signal, 1)
	dep.Log = logger.NewLogger()
	dep.Cfg, err = configs.New()
	if err != nil {
		return
	}

	return
}

var infraDependencies = func(dep *dependencies) (err error) {
	dep.DB, err = initializePostgres(dep.Cfg.DB)
	if err != nil {
		return
	}
	dep.Gin = initializeGin(dep.Log)

	return
}

var workerDependencies = func(dep *dependencies) (err error) {
	//dep.Workers.VisitWorker = workers.NewVisitWorker(dep.Log, dep.Repositories.Visit, dep.Repositories.Link)

	return
}

var serviceDependencies = func(dep *dependencies) (err error) {
	//dep.Services.User, err = userService.New(dep.Log, dep.Cfg.UserService, dep.Repositories.User, dep.Cache.VerificationCodeCache)
	//if err != nil {
	//	return
	//}
	//
	//dep.Services.Link = linkService.New(dep.Log, dep.Cfg.LinkService, dep.Repositories.Link, dep.Repositories.Visit, dep.Cache.LinkCache)
	return
}

var routersDependencies = func(dep *dependencies) (err error) {
	//dep.Routers.Middlewares.PublicAuth = middlewares.NewPublicAuthMiddleware(dep.Log)
	//dep.Routers.V1 = v1.New(dep.Controllers.AuthController, dep.Controllers.LinkController, dep.Controllers.DashboardController, dep.Controllers.UserController, dep.Routers.Middlewares.PublicAuth)
	//dep.Routers.Router = router.New(dep.Log, dep.Routers.V1, dep.Services.Link)
	return
}

var controllerDependencies = func(dep *dependencies) (err error) {
	//dep.Controllers.AuthController = authController.New(dep.Services.User)
	//dep.Controllers.LinkController = linkController.New(dep.Services.Link)
	//dep.Controllers.DashboardController = dashboardController.New(dep.Services.Link)
	//dep.Controllers.UserController = userController.New(dep.Services.User)
	return
}

var repositoryDependencies = func(dep *dependencies) (err error) {
	//dep.Repositories.User = userRepository.New(dep.Log, dep.DB)
	//dep.Repositories.Link = linkRepository.New(dep.Log, dep.DB)
	//dep.Repositories.Visit = visitRepository.New(dep.Log, dep.DB)
	return
}

func initializePostgres(configs gormext.Configs) (*gorm.DB, error) {
	dbInstance, err := db.NewDatabase(context.Background(), configs)
	if err != nil {
		return nil, err
	}

	pdb, err := dbInstance.DB.DB()
	if err != nil {
		log.WithError(err).Error("failed to get sql DB")
		return nil, err
	}
	pdb.SetConnMaxLifetime(time.Minute)
	pdb.SetMaxIdleConns(10)
	pdb.SetMaxOpenConns(200)

	if err = dbInstance.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx,
			gormext.UUIDExtension,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Error("failed to add extensions to database")
		return nil, err
	}

	return dbInstance.DB, nil
}

func initializeGin(log *log.Logger) *gin.Engine {
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	g := gin.Default()
	return g
}
