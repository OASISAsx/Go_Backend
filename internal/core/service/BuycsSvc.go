package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type buycsSvc struct {
	repo port.BuycsRepo
}

func NewBuycsSvc(repo port.BuycsRepo) domain.BuycsSvc {
	return buycsSvc{
		repo: repo,
	}
}

func (s buycsSvc) GetAllBuycs() ([]domain.BuycsRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get buycs form DB")
	}
	resp := []domain.BuycsRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BuycsRespone{
			RserveId:     c.RserveId,
			StoreId:      c.StoreId,
			UserId:       c.UserId,
			RecordStatus: c.RecordStatus,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s buycsSvc) GetBuycs(id int) (*domain.BuycsRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get buycs form DB")
	}
	resp := domain.BuycsRespone{
		RserveId:     cust.RserveId,
		StoreId:      cust.StoreId,
		UserId:       cust.UserId,
		RecordStatus: cust.RecordStatus,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}
func (s buycsSvc) AddBuycs(req domain.BuycsRequest) (*domain.BuycsRespone, error) {
	newtime := time.Now()
	cust := port.Buycs{
		StoreId:      req.StoreId,
		UserId:       req.UserId,
		RecordStatus: req.RecordStatus,
		CreatedBy:    req.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save buycs")
	}
	resp := domain.BuycsRespone{
		StoreId:      newCust.StoreId,
		UserId:       newCust.UserId,
		RecordStatus: newCust.RecordStatus,
		CreatedBy:    newCust.CreatedBy,
		CreatedDate:  newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s buycsSvc) UpdateBuycs(id int, req domain.BuycsRequest) error {
	newtime := time.Now()
	cust := port.Buycs{

		RecordStatus: req.RecordStatus,
		UpdatedBy:    req.UpdatedBy,
		UpdatedDate:  newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update buycs")
	}
	return nil
}
func (s buycsSvc) DeleteBuycs(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete buycs")
	}
	return nil
}
