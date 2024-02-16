package model

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)

type SortInfo struct {
	SortStrategy Order  `json:"sort_strategy"`
	SortBy       string `json:"sort_by"`
}

type PageInfo struct {
	Items      int64    `json:"items"`       //
	Offset     int64    `json:"offset"`      // The number of items to skip in the result set.
	PerPage    int64    `json:"per_page"`    // The maximum number of items to return per page.
	TotalPages int64    `json:"total_pages"` // The total number of pages available in the result set.
	SortInfo   SortInfo `json:"sort_info"`   // The order in which items are sorted.
}
