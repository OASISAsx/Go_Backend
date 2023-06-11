package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type areapvSvc struct {
	repo port.AreapvRepo
}

func NewAreapvSvc(repo port.AreapvRepo) domain.AreapvSvc {
	return areapvSvc{
		repo: repo,
	}
}

func (s areapvSvc) GetAllAreapv() ([]domain.AreapvRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areapv form DB")
	}
	resp := []domain.AreapvRespone{}
	for _, c := range custs {
		resp = append(resp, domain.AreapvRespone{
			ProvinceId:   c.ProvinceId,
			ProvinceName: c.ProvinceName,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s areapvSvc) GetAreapv(id int) (*domain.AreapvRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areapv form DB")
	}
	resp := domain.AreapvRespone{
		ProvinceId:   cust.ProvinceId,
		ProvinceName: cust.ProvinceName,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}
func (s areapvSvc) AddAreapv(req domain.AreapvRequest) (*domain.AreapvRespone, error) {
	newtime := time.Now()
	cust := port.Areapv{
		ProvinceName: req.ProvinceName,
		CreatedBy:    req.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save areapv")
	}
	resp := domain.AreapvRespone{
		ProvinceName: newCust.ProvinceName,
		CreatedBy:    newCust.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s areapvSvc) UpdateAreapv(id int, req domain.AreapvRequest) error {
	newtime := time.Now()
	cust := port.Areapv{

		ProvinceName: req.ProvinceName,
		UpdatedBy:    req.UpdatedBy,
		UpdatedDate:  newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update areapv")
	}
	return nil
}
func (s areapvSvc) DeleteAreapv(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete areapv")
	}
	return nil
}
