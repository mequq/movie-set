package app

import (
	"app/config"
	"app/internal/service"

	ginzerolog "github.com/easonlin404/gin-zerolog"
	"github.com/gin-gonic/gin"
)

// App is the application
type App struct {
	config  *config.Config
	Service *service.Service
	Engine  *gin.Engine
	Router  *gin.RouterGroup
}

// NewApp creates a new application.
func NewApp(config *config.Config, service *service.Service) *App {
	// create a new gin engine
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(ginzerolog.Logger())

	// create a new router
	router := engine.Group("/")
	// create a new app
	app := &App{
		config:  config,
		Service: service,
		Engine:  engine,
		Router:  router,
	}
	return app

}

// register routes
func (a *App) RegisterRoutes() {
	a.Service.RegisterRoutes(a.Router)
}

// Run runs the application.
func (a *App) Run() error {
	return a.Engine.Run(a.config.Server.Host + ":" + a.config.Server.Port)
}
