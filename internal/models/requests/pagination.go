package requests

type PaginationReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
