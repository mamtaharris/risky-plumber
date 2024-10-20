package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mamtaharris/risky-plumber/internal/models/entities"
	"github.com/mamtaharris/risky-plumber/internal/models/requests"
	"github.com/mamtaharris/risky-plumber/internal/models/responses"
)

func Test_riskService_Create(t *testing.T) {
	type fields struct {
		risks map[uuid.UUID]entities.Risks
	}
	type args struct {
		ctx     context.Context
		riskReq requests.RiskReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    responses.RiskResp
		wantErr bool
	}{
		{
			name: "happy case",
			fields: fields{
				risks: make(map[uuid.UUID]entities.Risks),
			},
			args: args{
				ctx: context.TODO(),
				riskReq: requests.RiskReq{
					State:       "open",
					Title:       "title",
					Description: "description",
				},
			},
			want: responses.RiskResp{
				State:       "open",
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &riskService{
				risks: tt.fields.risks,
			}
			got, err := r.Create(tt.args.ctx, tt.args.riskReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.State != tt.want.State || got.Title != tt.want.Title || got.Description != tt.want.Description {
				t.Errorf("riskService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_riskService_GetByID(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    responses.RiskResp
		wantErr bool
	}{
		{
			name: "happy case",
			want: responses.RiskResp{
				State:       "open",
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := uuid.New()
			r := &riskService{
				risks: map[uuid.UUID]entities.Risks{
					id: {
						ID:          id,
						State:       "open",
						Title:       "title",
						Description: "description",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
			}
			got, err := r.GetByID(tt.args.ctx, id)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.State != tt.want.State || got.Title != tt.want.Title || got.Description != tt.want.Description {
				t.Errorf("riskService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_riskService_GetAll(t *testing.T) {
	type fields struct {
		risks map[uuid.UUID]entities.Risks
	}
	type args struct {
		ctx           context.Context
		paginationReq requests.PaginationReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []responses.RiskResp
		wantErr bool
	}{
		{
			name: "GetAll - return all risks without pagination",
			fields: fields{
				risks: map[uuid.UUID]entities.Risks{
					uuid.New(): {
						ID:          uuid.New(),
						State:       "open",
						Title:       "Risk 1",
						Description: "First risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
					uuid.New(): {
						ID:          uuid.New(),
						State:       "closed",
						Title:       "Risk 2",
						Description: "Second risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				paginationReq: requests.PaginationReq{
					Offset: 0,
					Limit:  10,
				},
			},
			want: []responses.RiskResp{
				{
					State:       "open",
					Title:       "Risk 1",
					Description: "First risk",
				},
				{
					State:       "closed",
					Title:       "Risk 2",
					Description: "Second risk",
				},
			},
			wantErr: false,
		},
		{
			name: "GetAll - return paginated risks",
			fields: fields{
				risks: map[uuid.UUID]entities.Risks{
					uuid.New(): {
						ID:          uuid.New(),
						State:       "open",
						Title:       "Risk 1",
						Description: "First risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
					uuid.New(): {
						ID:          uuid.New(),
						State:       "closed",
						Title:       "Risk 2",
						Description: "Second risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
					uuid.New(): {
						ID:          uuid.New(),
						State:       "in-progress",
						Title:       "Risk 3",
						Description: "Third risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				paginationReq: requests.PaginationReq{
					Offset: 1,
					Limit:  1,
				},
			},
			want: []responses.RiskResp{
				{
					State:       "closed",
					Title:       "Risk 2",
					Description: "Second risk",
				},
			},
			wantErr: false,
		},
		{
			name: "GetAll - offset greater than number of risks",
			fields: fields{
				risks: map[uuid.UUID]entities.Risks{
					uuid.New(): {
						ID:          uuid.New(),
						State:       "open",
						Title:       "Risk 1",
						Description: "First risk",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				paginationReq: requests.PaginationReq{
					Offset: 10,
					Limit:  10,
				},
			},
			want:    []responses.RiskResp{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &riskService{
				risks: tt.fields.risks,
			}
			got, err := r.GetAll(tt.args.ctx, tt.args.paginationReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("riskService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Compare the number of risks returned
				if len(got) != len(tt.want) {
					t.Errorf("riskService.GetAll() returned %d risks, want %d", len(got), len(tt.want))
				}

				// Compare the static fields in each risk
				for i := range got {
					if got[i].State != tt.want[i].State || got[i].Title != tt.want[i].Title || got[i].Description != tt.want[i].Description {
						t.Errorf("riskService.GetAll() = %v, want %v", got[i], tt.want[i])
					}
				}
			}
		})
	}
}
