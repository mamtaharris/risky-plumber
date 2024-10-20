package validators

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mamtaharris/risky-plumber/config"
	"github.com/mamtaharris/risky-plumber/internal/models/requests"
	vMock "github.com/mamtaharris/risky-plumber/internal/validators/mocks"
)

func Test_riskReqValidator_ValidateCreateRiskReq(t *testing.T) {
	type fields struct {
		validator ValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator func(ctrl *gomock.Controller) *vMock.MockValidatorInterface
		want          requests.RiskReq
		wantErr       bool
	}{
		{
			name: "unknown params present",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"test": "test"
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(errors.New("dummy"))
				return repaymentMock
			},
			want:    requests.RiskReq{},
			wantErr: true,
		},
		{
			name: "invalid request",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"amount": 100
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RiskReq{},
			wantErr: true,
		},
		{
			name: "happy case",
			args: args{ctx: &gin.Context{Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"state": "open"
			}`)))}},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				repaymentMock := vMock.NewMockValidatorInterface(ctrl)
				repaymentMock.EXPECT().ValidateUnknownParams(gomock.Any(), gomock.Any()).Return(nil)
				return repaymentMock
			},
			want:    requests.RiskReq{State: "open"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewRiskValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateCreateRiskReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskReqValidator.ValidateCreateRiskReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("riskReqValidator.ValidateCreateRiskReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_riskReqValidator_ValidateGetRiskReq(t *testing.T) {
	type fields struct {
		validator ValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator func(ctrl *gomock.Controller) *vMock.MockValidatorInterface
		want          uuid.UUID
		wantErr       bool
	}{
		{
			name: "invalid UUID format",
			args: args{ctx: &gin.Context{
				Params: gin.Params{
					{Key: "id", Value: "invalid-uuid"},
				},
			}},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want:    uuid.Nil,
			wantErr: true,
		},
		{
			name: "valid UUID format",
			args: args{ctx: &gin.Context{
				Params: gin.Params{
					{Key: "id", Value: "123e4567-e89b-12d3-a456-426614174000"},
				},
			}},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want:    uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewRiskValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateGetRiskReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskReqValidator.ValidateGetRiskReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("riskReqValidator.ValidateGetRiskReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_riskReqValidator_ValidateGetAllRiskReq(t *testing.T) {
	type fields struct {
		validator ValidatorInterface
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator func(ctrl *gomock.Controller) *vMock.MockValidatorInterface
		want          requests.PaginationReq
		wantErr       bool
	}{
		{
			name: "default pagination settings",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						URL: &url.URL{RawQuery: ""},
					},
				},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want: requests.PaginationReq{
				Limit:  10,
				Offset: 0,
			},
			wantErr: false,
		},
		{
			name: "valid limit and offset from query params",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						URL: &url.URL{RawQuery: "limit=20&offset=10"},
					},
				},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want: requests.PaginationReq{
				Limit:  20,
				Offset: 10,
			},
			wantErr: false,
		},
		{
			name: "invalid limit, fallback to default",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						URL: &url.URL{RawQuery: "limit=invalid&offset=5"},
					},
				},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want: requests.PaginationReq{
				Limit:  10,
				Offset: 5,
			},
			wantErr: false,
		},
		{
			name: "invalid offset, fallback to default",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						URL: &url.URL{RawQuery: "limit=15&offset=invalid"},
					},
				},
			},
			mockValidator: func(ctrl *gomock.Controller) *vMock.MockValidatorInterface {
				return vMock.NewMockValidatorInterface(ctrl)
			},
			want: requests.PaginationReq{
				Limit:  15,
				Offset: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitConfig()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.validator = tt.mockValidator(ctrl)

			v := NewRiskValidator(
				tt.fields.validator,
			)
			got, err := v.ValidateGetAllRiskReq(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskReqValidator.ValidateGetAllRiskReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("riskReqValidator.ValidateGetAllRiskReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
