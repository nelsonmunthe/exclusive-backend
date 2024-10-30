package dto

type QUeryRequest struct {
	Pagination PaginationRequest
	Search     string `form:"q" json:"q"`
}
