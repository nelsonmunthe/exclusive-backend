package dto

type PaginationRequest struct {
	Page    int `form:"page" json:"page"`
	PerPage int `form:"perPage" json:"perPage"`
}
