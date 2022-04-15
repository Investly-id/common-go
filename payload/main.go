package payload

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"messages"`
	Meta    *Pagination `json:"meta,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
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
