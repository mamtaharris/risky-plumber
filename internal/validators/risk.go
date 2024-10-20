package validators

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mamtaharris/risky-plumber/config"
	"github.com/mamtaharris/risky-plumber/internal/constants"
	"github.com/mamtaharris/risky-plumber/internal/models/requests"
)

type riskReqValidator struct {
	validator ValidatorInterface
}

func NewRiskValidator(validator ValidatorInterface) RiskReqValidatorInterface {
	return &riskReqValidator{
		validator: validator,
	}
}

//go:generate mockgen -package mocks -source=risk.go -destination=mocks/risk_mocks.go
type RiskReqValidatorInterface interface {
	ValidateCreateRiskReq(ctx *gin.Context) (requests.RiskReq, error)
	ValidateGetRiskReq(ctx *gin.Context) (uuid.UUID, error)
	ValidateGetAllRiskReq(ctx *gin.Context) (requests.PaginationReq, error)
}

func (v *riskReqValidator) ValidateCreateRiskReq(ctx *gin.Context) (requests.RiskReq, error) {
	var reqBody requests.RiskReq
	err := v.validator.ValidateUnknownParams(&reqBody, ctx)
	if err != nil {
		return requests.RiskReq{}, err
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return requests.RiskReq{}, err
	}
	validStates := map[string]bool{
		constants.RiskStates.Open:          true,
		constants.RiskStates.Closed:        true,
		constants.RiskStates.Accepted:      true,
		constants.RiskStates.Investigating: true,
	}
	if !validStates[reqBody.State] {
		return requests.RiskReq{}, errors.New("invalid state")
	}

	return reqBody, nil
}

func (v *riskReqValidator) ValidateGetRiskReq(ctx *gin.Context) (uuid.UUID, error) {
	riskID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return uuid.Nil, err
	}
	return riskID, nil
}

func (v *riskReqValidator) ValidateGetAllRiskReq(ctx *gin.Context) (requests.PaginationReq, error) {
	reqBody := requests.PaginationReq{
		Limit:  config.Pagination.Limit,
		Offset: config.Pagination.Offset,
	}

	if val := ctx.Query("limit"); val != "" {
		parsedLimit, err := strconv.Atoi(val)
		if err == nil && parsedLimit > 0 {
			reqBody.Limit = parsedLimit
		}
	}
	if val := ctx.Query("offset"); val != "" {
		parsedOffset, err := strconv.Atoi(val)
		if err == nil && parsedOffset > 0 {
			reqBody.Offset = parsedOffset
		}
	}

	return reqBody, nil
}
