package dto

type PaginationResponse struct {
	PerPage      int  `json:"perPage"`
	Total        uint `json:"total"`
	CurrentPage  int  `json:"currentPage"`
	PreviousPage int  `json:"previousPage"`
	NextPage     int  `json:"nextPage"`
}

func (p *PaginationResponse) Evaluate() {
	if p.CurrentPage-1 > 0 {
		p.PreviousPage = p.CurrentPage - 1
	}

	if uint(p.CurrentPage*p.PerPage) < p.Total {
		p.NextPage = p.CurrentPage + 1
	}
}
