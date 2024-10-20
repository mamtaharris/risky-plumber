package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mamtaharris/risky-plumber/internal/models/entities"
	"github.com/mamtaharris/risky-plumber/internal/models/requests"
	"github.com/mamtaharris/risky-plumber/internal/models/responses"
)

type riskService struct {
	risks map[uuid.UUID]entities.Risks
}

func NewRiskService() RiskService {
	return &riskService{
		risks: make(map[uuid.UUID]entities.Risks),
	}
}

//go:generate mockgen -package mocks -source=risk.go -destination=mocks/risk_mocks.go
type RiskService interface {
	Create(ctx context.Context, riskReq requests.RiskReq) (responses.RiskResp, error)
	GetByID(ctx context.Context, id uuid.UUID) (responses.RiskResp, error)
	GetAll(ctx context.Context, paginationReq requests.PaginationReq) ([]responses.RiskResp, error)
}

func (r *riskService) Create(ctx context.Context, riskReq requests.RiskReq) (responses.RiskResp, error) {
	risk := entities.Risks{
		ID:          uuid.New(),
		State:       riskReq.State,
		Title:       riskReq.Title,
		Description: riskReq.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	r.risks[risk.ID] = risk

	riskResp := responses.RiskResp{
		ID:          risk.ID,
		State:       risk.State,
		Title:       risk.Title,
		Description: risk.Description,
		CreatedAt:   int(risk.CreatedAt.Unix()),
		UpdatedAt:   int(risk.UpdatedAt.Unix()),
	}

	return riskResp, nil
}

func (r *riskService) GetByID(ctx context.Context, id uuid.UUID) (responses.RiskResp, error) {
	risk, exists := r.risks[id]
	if !exists {
		return responses.RiskResp{}, errors.New("risk with id not found: " + id.String())
	}

	return responses.RiskResp{
		ID:          risk.ID,
		State:       risk.State,
		Title:       risk.Title,
		Description: risk.Description,
		CreatedAt:   int(risk.CreatedAt.Unix()),
		UpdatedAt:   int(risk.UpdatedAt.Unix()),
	}, nil
}

func (r *riskService) GetAll(ctx context.Context, paginationReq requests.PaginationReq) ([]responses.RiskResp, error) {
	var allRisks []responses.RiskResp
	for _, risk := range r.risks {
		allRisks = append(allRisks, responses.RiskResp{
			ID:          risk.ID,
			State:       risk.State,
			Title:       risk.Title,
			Description: risk.Description,
			CreatedAt:   int(risk.CreatedAt.Unix()),
			UpdatedAt:   int(risk.UpdatedAt.Unix()),
		})
	}

	// Handle offset greater than the number of risks
	if paginationReq.Offset >= len(allRisks) {
		return []responses.RiskResp{}, nil
	}

	// Paginate the results
	end := paginationReq.Offset + paginationReq.Limit
	if end > len(allRisks) {
		end = len(allRisks)
	}
	paginatedRisks := allRisks[paginationReq.Offset:end]

	return paginatedRisks, nil
}
