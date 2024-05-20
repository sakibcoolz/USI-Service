package urlmapping

import (
	"USI-Service/config"
	"USI-Service/controller"
	dbm "USI-Service/db"
	"USI-Service/domain"
	"USI-Service/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Server(config config.IConfig, log *zap.Logger) {

	// connect to database
	db, err := dbm.Connect(config, log)
	if err != nil {
		log.Fatal("Error while connecting to database", zap.Error(err))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	// take database connection and pass it to domain
	dbs := domain.NewDomain(db, log)

	dbm.DBMigrate(db)

	srv := service.NewService(dbs, log)

	controller := controller.NewController(srv, log, validate)

	go TerminateService(StopService, log)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"*"},
	// 	AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	// 	AllowHeaders:  []string{"Origin"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// 	// AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	return origin == "https://github.com"
	// 	// },
	// 	MaxAge: 12 * time.Hour,
	// }))

	apiRigRouter := router.Group("/register/")
	apiAuthRouter := router.Group("/api/")

	apiAuthRouter.Use(func(c *gin.Context) {
		c.Next()
	})

	apiRigRouter.POST("/register", controller.Register)

	apiAuthRouter.POST("/login", controller.Login)

	apiAuthRouter.GET("/userdeatils/:user", controller.UserDetails)

	if err := router.Run(config.GetService().Host + ":" + config.GetService().Port); err != nil {
		log.Fatal("Error while starting server", zap.Error(err))
	}
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	select {
	case <-stopService:
		log.Info("Terminating service")

		os.Exit(0)
	}
}
