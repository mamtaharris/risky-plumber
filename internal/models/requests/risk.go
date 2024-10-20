package requests

type RiskReq struct {
	State       string `json:"state" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
