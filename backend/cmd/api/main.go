package main

import (
	"net/http"

	"github.com/example/entrax/backend/internal/config"
	"github.com/example/entrax/backend/internal/handlers"
	"github.com/example/entrax/backend/internal/middleware"
	"github.com/example/entrax/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogger(log))
	router.Use(middleware.ErrorHandler(log))

	healthHandler := handlers.NewHealthHandler()
	projectHandler := handlers.NewProjectHandler()
	authMiddleware := middleware.NewAzureADJWTMiddleware(cfg)

	router.GET("/health", healthHandler.GetHealth)

	api := router.Group("/api/v1")
	{
		api.GET("/projects", authMiddleware.ValidateToken(), projectHandler.ListProjects)
		api.POST("/projects", authMiddleware.ValidateToken(), projectHandler.CreateProject)
	}

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Info().Str("port", cfg.Port).Msg("starting API server")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server terminated unexpectedly")
	}
}
