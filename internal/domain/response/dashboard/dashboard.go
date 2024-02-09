package dashboard

type DashboardResponse struct {
	User            int   `json:"user"`
	Product         int   `json:"product"`
	Order           int   `json:"order"`
	Pendapatan      int   `json:"pendapatan"`
	TotalPendapatan []int `json:"total_pendapatan"`
}
