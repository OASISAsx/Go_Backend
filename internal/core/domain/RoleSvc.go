package domain

type RoleSvc interface {
	GetAllRole() ([]RoleResp, error)
	GetRole(int) (*RoleResp, error)
	AddRole(RoleReq) (*RoleResp, error)
	UpdateRole(int, RoleReq) error
	DeleteRole(int) error
}

type RoleReq struct {
	RoleID      uint   `json:"roleid"`
	RoleName    string `json:"rolename"`
	RoleDesc    string `json:"roledesc"`
	Status      string `json:"status"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type RoleResp struct {
	RoleID      uint   `json:"roleid"`
	RoleName    string `json:"rolename"`
	RoleDesc    string `json:"roledesc"`
	Status      string `json:"status"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
