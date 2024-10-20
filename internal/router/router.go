package router

import (
	"context"

	"github.com/gin-gonic/gin"
	handler "github.com/mamtaharris/risky-plumber/internal/handlers"

	"github.com/mamtaharris/risky-plumber/internal/services"

	"github.com/mamtaharris/risky-plumber/internal/validators"
)

func SetRouter(ctx context.Context) (*gin.Engine, error) {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	riskSvc := services.NewRiskService()
	validator := validators.NewValidator()
	riskValidator := validators.NewRiskValidator(validator)
	riskHandler := handler.NewRiskHandler(riskSvc, riskValidator)

	routerV1 := router.Group("/v1/risks")
	routerV1.POST("/", riskHandler.Create)
	routerV1.GET("/:id", riskHandler.GetByID)
	routerV1.GET("/", riskHandler.GetAll)

	return router, nil
}
