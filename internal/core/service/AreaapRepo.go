package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type areaapSvc struct {
	repo port.AreaapRepo
}

func NewAreaapSvc(repo port.AreaapRepo) domain.AreaapSvc {
	return areaapSvc{
		repo: repo,
	}
}

func (s areaapSvc) GetAllAreaap() ([]domain.AreaapRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areaap form DB")
	}
	resp := []domain.AreaapRespone{}
	for _, c := range custs {
		resp = append(resp, domain.AreaapRespone{
			AmphurId:    c.AmphurId,
			AmphurName:  c.AmphurName,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s areaapSvc) GetAreaap(id int) (*domain.AreaapRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get areaap form DB")
	}
	resp := domain.AreaapRespone{
		AmphurId:    cust.AmphurId,
		AmphurName:  cust.AmphurName,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}
func (s areaapSvc) AddAreaap(req domain.AreaapRequest) (*domain.AreaapRespone, error) {
	newtime := time.Now()
	cust := port.Areaap{
		AmphurName:  req.AmphurName,
		CreatedBy:   req.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save areaap")
	}
	resp := domain.AreaapRespone{
		AmphurName:  newCust.AmphurName,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s areaapSvc) UpdateAreaap(id int, req domain.AreaapRequest) error {
	newtime := time.Now()
	cust := port.Areaap{

		AmphurName:  req.AmphurName,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update areaap")
	}
	return nil
}

func (s areaapSvc) DeleteAreaap(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete areaap")
	}
	return nil
}
