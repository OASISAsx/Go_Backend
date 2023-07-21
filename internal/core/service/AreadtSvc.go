package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type areadtSvc struct {
	repo port.AreadtRepo
}

func NewAreadtSvc(repo port.AreadtRepo) domain.AreadtSvc {
	return areadtSvc{
		repo: repo,
	}
}

func (s areadtSvc) GetAllAreadt() ([]domain.AreadtRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areadt form DB")
	}
	resp := []domain.AreadtRespone{}
	for _, c := range custs {
		resp = append(resp, domain.AreadtRespone{
			DistrictId:   c.DistrictId,
			DistrictCode: c.DistrictCode,
			DistrictName: c.DistrictName,
			RefAmphurId:  c.RefAmphurId,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s areadtSvc) GetAreadt(id int) (*domain.AreadtRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areadt form DB")
	}
	resp := domain.AreadtRespone{
		DistrictId:   cust.DistrictId,
		DistrictCode: cust.DistrictCode,
		DistrictName: cust.DistrictName,
		RefAmphurId:  cust.RefAmphurId,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}
func (s areadtSvc) AddAreadt(req domain.AreadtRequest) (*domain.AreadtRespone, error) {
	newtime := time.Now()
	cust := port.Areadt{
		DistrictName: req.DistrictName,
		RefAmphurId:  req.RefAmphurId,
		CreatedBy:    req.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save areadt")
	}
	resp := domain.AreadtRespone{
		DistrictName: newCust.DistrictName,
		RefAmphurId:  newCust.RefAmphurId,
		CreatedBy:    newCust.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s areadtSvc) UpdateAreadt(id int, req domain.AreadtRequest) error {
	newtime := time.Now()
	cust := port.Areadt{

		DistrictName: req.DistrictName,
		RefAmphurId:  req.RefAmphurId,
		UpdatedBy:    req.UpdatedBy,
		UpdatedDate:  newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update areadt")
	}
	return nil
}
func (s areadtSvc) DeleteAreadt(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete areadt")
	}
	return nil
}
