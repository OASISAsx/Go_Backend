package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type roleSvc struct {
	repo port.RoleRepo
}

func NewRoleSvc(repo port.RoleRepo) domain.RoleSvc {
	return roleSvc{
		repo: repo,
	}
}

func (s roleSvc) GetAllRole() ([]domain.RoleResp, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get role form DB")
	}
	resp := []domain.RoleResp{}
	for _, c := range custs {
		resp = append(resp, domain.RoleResp{
			RoleID:      c.RoleID,
			RoleName:    c.RoleName,
			RoleDesc:    c.RoleDesc,
			Status:      c.Status,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s roleSvc) GetRole(id int) (*domain.RoleResp, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get role form DB")
	}
	resp := domain.RoleResp{
		RoleID:      cust.RoleID,
		RoleName:    cust.RoleName,
		RoleDesc:    cust.RoleDesc,
		Status:      cust.Status,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}
func (s roleSvc) AddRole(req domain.RoleReq) (*domain.RoleResp, error) {
	newtime := time.Now()
	cust := port.Role{
		RoleName:    req.RoleName,
		RoleDesc:    req.RoleDesc,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save role")
	}
	resp := domain.RoleResp{
		RoleID:      newCust.RoleID,
		RoleName:    newCust.RoleName,
		RoleDesc:    newCust.RoleDesc,
		Status:      newCust.Status,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s roleSvc) UpdateRole(id int, req domain.RoleReq) error {
	newtime := time.Now()
	cust := port.Role{
		RoleName:    req.RoleName,
		RoleDesc:    req.RoleDesc,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update role")
	}
	return nil
}
func (s roleSvc) DeleteRole(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete role")
	}
	return nil
}
