package domain

type ReportSvc interface {
	GetAllReport() ([]ReportResp, error)
	GetReport(int) (*ReportResp, error)
	AddReport(ReportReq) (*ReportResp, error)
	UpdateReport(int, ReportReq) error
	DeleteReport(int) error
}

type ReportReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

type ReportResp struct {
	ReportId uint   `json:"reportid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}
