package resource

// PaginationQueryParam is a pagination query param
type PaginationQueryParam struct {
	Query    string `form:"query" json:"query"`
	Sort     string `form:"sort" json:"sort"`
	Order    string `form:"order" json:"order"`
	Limit    int    `form:"limit,default=10" json:"limit"`
	Offset   int    `form:"offset,default=0" json:"offset"`
	GameCode string `form:"game_code" json:"game_code"`
	Slug     string `form:"slug" json:"slug"`
	Lang     string `form:"lang" json:"lang"`
	Status   string `form:"status" json:"status"`
}

// Meta is a meta response
type Meta struct {
	Total_Data   int64 `json:"total_data"`
	Per_Page     int   `json:"per_page"`
	Current_Page int   `json:"current_page"`
	Total_Page   int64 `json:"total_page"`
}

type DashboardMeta struct {
	Total_Buku_Tersedia int64 `json:"total_buku_tersedia"`
	Total_Buku_Dipinjam int64 `json:"total_buku_dipinjam"`
	Total_User          int64 `json:"total_user"`
	Total_User_pinjam   int64 `json:"total_user_pinjam"`
}
