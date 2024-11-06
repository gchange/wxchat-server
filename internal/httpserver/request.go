package httpserver

type PageRequest struct {
	Page       int    `form:"page" json:"page"`
	PageSize   int    `form:"page_size" json:"page_size"`
	Sort       string `form:"sort" json:"sort"`
	Descending bool   `form:"descending" json:"descending"`
}
