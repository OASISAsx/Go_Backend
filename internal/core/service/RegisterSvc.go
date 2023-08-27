package service

import (
	"collection/errs"
	"collection/infrastructure"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"collection/utils"
	"errors"
	"net/http"
	"time"
)

type registerSvc struct {
	repo port.RegisterRepo
}

func NewRegisterSvc(repo port.RegisterRepo) domain.RegisterSvc {
	return registerSvc{
		repo: repo,
	}
}

func (s registerSvc) GetAllRegister() ([]domain.RegisterResp, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get register form DB")
	}
	resp := []domain.RegisterResp{}
	for _, c := range custs {
		resp = append(resp, domain.RegisterResp{
			UserId:       c.UserId,
			RoleId:       c.RoleId,
			Username:     c.Username,
			Password:     c.Password,
			Nickname:     c.Nickname,
			Email:        c.Email,
			RecordStatus: c.RecordStatus,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s registerSvc) GetRegister(id int) (*domain.RegisterResp, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get register form DB")
	}
	resp := domain.RegisterResp{
		UserId:       cust.UserId,
		RoleId:       cust.RoleId,
		Username:     cust.Username,
		Password:     cust.Password,
		Nickname:     cust.Nickname,
		Email:        cust.Email,
		RecordStatus: cust.RecordStatus,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}
func (s registerSvc) AddRegister(req domain.RegisterReq) (*domain.RegisterResp, error) {
	newtime := time.Now()
	hashpwd, _ := utils.HashPassword(req.Password)
	cust := port.Register{
		RoleId:       "user",
		Username:     req.Username,
		Password:     hashpwd,
		Nickname:     req.Nickname,
		Email:        req.Email,
		RecordStatus: true,
		CreatedBy:    "System",
		CreatedDate:  newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save register")
	}
	resp := domain.RegisterResp{
		UserId:       newCust.UserId,
		RoleId:       newCust.RoleId,
		Username:     newCust.Username,
		Password:     newCust.Password,
		Nickname:     newCust.Nickname,
		Email:        newCust.Email,
		RecordStatus: newCust.RecordStatus,
		CreatedBy:    newCust.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s registerSvc) UpdateRegister(id int, req domain.RegisterReq) error {
	newtime := time.Now()
	hashpwd, _ := utils.HashPassword(req.Password)
	cust := port.Register{

		Password:    hashpwd,
		Nickname:    req.Nickname,
		Email:       req.Email,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update register")
	}
	return nil
}
func (s registerSvc) UpdateRole(id int, req domain.RoleId) error {
	newtime := time.Now()
	cust := port.Register{

		RoleId:      req.RoleId,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.UpdateRole(id, cust.RoleId)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update register")
	}
	return nil
}
func (s registerSvc) UpdateStetus(id int, req domain.RecordStatus) error {
	newtime := time.Now()
	cust := port.Register{

		RecordStatus: req.RecordStatus,
		UpdatedDate:  newtime.Format(time.DateTime),
	}
	err := s.repo.UpdateStatus(id, cust.RecordStatus)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update register")
	}
	return nil
}

func (s registerSvc) DeleteRegister(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete register")
	}
	return nil
}

// func (s registerSvc) FindByUsername(username string) (Register, error){

// }
func (s registerSvc) GetProfile(username string) (*domain.RegisterResp, error) {

	cust, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get register form DB")
	}
	resp := domain.RegisterResp{
		UserId:       cust.UserId,
		RoleId:       cust.RoleId,
		Username:     cust.Username,
		Password:     cust.Password,
		Nickname:     cust.Nickname,
		Email:        cust.Email,
		RecordStatus: cust.RecordStatus,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}
func (s registerSvc) Login(users domain.LoginReq) (string, error) {
	// Find username in database
	new_users, users_err := s.repo.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := infrastructure.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.UserId, config.TokenSecret)
	errs.ErrorPanic(err_token)
	return token, nil
}
