package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamtaharris/risky-plumber/internal/models/responses"
	"github.com/mamtaharris/risky-plumber/internal/services"
	"github.com/mamtaharris/risky-plumber/internal/validators"
)

type RiskHandler struct {
	riskService   services.RiskService
	riskValidator validators.RiskReqValidatorInterface
}

func NewRiskHandler(riskService services.RiskService, riskValidator validators.RiskReqValidatorInterface) *RiskHandler {
	return &RiskHandler{riskService: riskService, riskValidator: riskValidator}
}

func (h *RiskHandler) Create(ctx *gin.Context) {
	req, err := h.riskValidator.ValidateCreateRiskReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.riskService.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *RiskHandler) GetByID(ctx *gin.Context) {
	reqID, err := h.riskValidator.ValidateGetRiskReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.riskService.GetByID(ctx, reqID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *RiskHandler) GetAll(ctx *gin.Context) {
	req, err := h.riskValidator.ValidateGetAllRiskReq(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResp{Error: err.Error()})
		return
	}
	response, err := h.riskService.GetAll(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResp{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
