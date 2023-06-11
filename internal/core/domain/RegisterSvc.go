package domain

type RegisterSvc interface {
	Login(users LoginReq) (string, error)
	GetProfile(string) (*RegisterResp, error)
	GetAllRegister() ([]RegisterResp, error)
	GetRegister(int) (*RegisterResp, error)
	AddRegister(RegisterReq) (*RegisterResp, error)
	UpdateRegister(int, RegisterReq) error
	DeleteRegister(int) error
}

type RegisterReq struct {
	RoleId       int    `json:"roleId" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Nickname     string `json:"nickname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	RecordStatus string `json:"recordstatus" binding:"required"`
	CreatedBy    string `json:"createdby" binding:"required"`
	CreatedDate  string `json:"createddate" binding:"required"`
	UpdatedBy    string `json:"updatedby" binding:"required"`
	UpdatedDate  string `json:"updateddate" binding:"required"`
}

type RegisterResp struct {
	UserId       uint   `json:"userid" binding:"required"`
	RoleId       int    `json:"roleId" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Nickname     string `json:"nickname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	RecordStatus string `json:"recordstatus" binding:"required"`
	CreatedBy    string `json:"createdby" binding:"required"`
	CreatedDate  string `json:"createddate" binding:"required"`
	UpdatedBy    string `json:"updatedby" binding:"required"`
	UpdatedDate  string `json:"updateddate" binding:"required"`
}
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResp struct {
	TokenType string      `json:"tokentype" binding:"required"`
	Token     string      `json:"token" binding:"required"`
	Message   string      `json:"message" binding:"required"`
	User      interface{} `json:"user,omitempty"`
	Bearer    interface{} `json:"Bearer,omitempty"`
}
type Resp struct {
	Status  string      `json:"stetus" binding:"required"`
	Code    int         `json:"code" binding:"required"`
	Message string      `json:"message" binding:"required"`
	User    interface{} `json:"user,omitempty"`
	Token   interface{} `json:"token,omitempty"`
}
