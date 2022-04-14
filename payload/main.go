package payload

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  bool   `json:"status"`
}

type Pagination struct {
	PerPage     int64 `json:"per_page"`
	CurrentPage int64 `json:"current_page"`
	LastPage    int64 `json:"last_page"`
	IsLoadMore  bool  `json:"is_load_more"`
}

type PaginationRequest struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"perpage"`
	Limit   int64
	Offset  int64
}
