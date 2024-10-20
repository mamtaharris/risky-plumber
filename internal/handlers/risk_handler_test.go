package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mamtaharris/risky-plumber/internal/models/requests"
	"github.com/mamtaharris/risky-plumber/internal/models/responses"
	"github.com/mamtaharris/risky-plumber/internal/services"
	sMock "github.com/mamtaharris/risky-plumber/internal/services/mocks"
	"github.com/mamtaharris/risky-plumber/internal/validators"
	vMock "github.com/mamtaharris/risky-plumber/internal/validators/mocks"
)

func TestRiskHandler(t *testing.T) {
}

func TestRiskHandler_Create(t *testing.T) {
	type fields struct {
		riskService   services.RiskService
		riskValidator validators.RiskReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		mockRiskService   func(ctrl *gomock.Controller) *sMock.MockRiskService
		mockRiskValidator func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface
		args              args
	}{
		{
			name: "failed at service layer",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(responses.RiskResp{}, errors.New("dummy"))
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateCreateRiskReq(gomock.Any()).Return(requests.RiskReq{}, nil)
				return riskMock
			},
		},
		{
			name: "failed at validator",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateCreateRiskReq(gomock.Any()).Return(requests.RiskReq{}, errors.New("dummy"))
				return riskMock
			},
		},
		{
			name: "happy case",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(responses.RiskResp{}, nil)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateCreateRiskReq(gomock.Any()).Return(requests.RiskReq{}, nil)
				return riskMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.riskService = tt.mockRiskService(ctrl)
			tt.fields.riskValidator = tt.mockRiskValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewRiskHandler(
				tt.fields.riskService,
				tt.fields.riskValidator,
			)
			h.Create(tt.args.ctx)
		})
	}
}

func TestRiskHandler_GetByID(t *testing.T) {
	type fields struct {
		riskService   services.RiskService
		riskValidator validators.RiskReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		mockRiskService   func(ctrl *gomock.Controller) *sMock.MockRiskService
		mockRiskValidator func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface
		args              args
	}{
		{
			name: "failed at service layer",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(responses.RiskResp{}, errors.New("dummy"))
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetRiskReq(gomock.Any()).Return(uuid.New(), nil)
				return riskMock
			},
		},
		{
			name: "failed at validator",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetRiskReq(gomock.Any()).Return(uuid.Nil, errors.New("dummy"))
				return riskMock
			},
		},
		{
			name: "happy case",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(responses.RiskResp{}, nil)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetRiskReq(gomock.Any()).Return(uuid.New(), nil)
				return riskMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.riskService = tt.mockRiskService(ctrl)
			tt.fields.riskValidator = tt.mockRiskValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewRiskHandler(
				tt.fields.riskService,
				tt.fields.riskValidator,
			)
			h.GetByID(tt.args.ctx)
		})
	}
}

func TestRiskHandler_GetAll(t *testing.T) {
	type fields struct {
		riskService   services.RiskService
		riskValidator validators.RiskReqValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		mockRiskService   func(ctrl *gomock.Controller) *sMock.MockRiskService
		mockRiskValidator func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface
		args              args
	}{
		{
			name: "failed at service layer",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return([]responses.RiskResp{}, errors.New("dummy"))
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetAllRiskReq(gomock.Any()).Return(requests.PaginationReq{}, nil)
				return riskMock
			},
		},
		{
			name: "failed at validator",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetAllRiskReq(gomock.Any()).Return(requests.PaginationReq{}, errors.New("dummy"))
				return riskMock
			},
		},
		{
			name: "happy case",
			mockRiskService: func(ctrl *gomock.Controller) *sMock.MockRiskService {
				riskMock := sMock.NewMockRiskService(ctrl)
				riskMock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return([]responses.RiskResp{}, nil)
				return riskMock
			},
			mockRiskValidator: func(ctrl *gomock.Controller) *vMock.MockRiskReqValidatorInterface {
				riskMock := vMock.NewMockRiskReqValidatorInterface(ctrl)
				riskMock.EXPECT().ValidateGetAllRiskReq(gomock.Any()).Return(requests.PaginationReq{}, nil)
				return riskMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.riskService = tt.mockRiskService(ctrl)
			tt.fields.riskValidator = tt.mockRiskValidator(ctrl)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.args.ctx = ctx

			h := NewRiskHandler(
				tt.fields.riskService,
				tt.fields.riskValidator,
			)
			h.GetAll(tt.args.ctx)
		})
	}
}
