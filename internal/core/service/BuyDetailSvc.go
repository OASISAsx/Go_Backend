package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type buydetailSvc struct {
	repo port.BuyDetailRepo
}

func NewBuyDetailSvc(repo port.BuyDetailRepo) domain.BuyDetailSvc {
	return buydetailSvc{
		repo: repo,
	}
}

func (s buydetailSvc) GetAllBuyDetail() ([]domain.BuyDetailRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get buydetail form DB")
	}
	resp := []domain.BuyDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BuyDetailRespone{
			BuyId:       c.BuyId,
			StoreId:     c.StoreId,
			ReserveId:   c.ReserveId,
			ProductId:   c.ProductId,
			QrCode:      c.QrCode,
			PayId:       c.PayId,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s buydetailSvc) GetBuyDetail(id int) (*domain.BuyDetailRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get buydetail form DB")
	}
	resp := domain.BuyDetailRespone{
		BuyId:       cust.BuyId,
		StoreId:     cust.StoreId,
		ReserveId:   cust.ReserveId,
		ProductId:   cust.ProductId,
		QrCode:      cust.QrCode,
		PayId:       cust.PayId,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}
func (s buydetailSvc) AddBuyDetail(req domain.BuyDetailRequest) (*domain.BuyDetailRespone, error) {
	newtime := time.Now()
	cust := port.BuyDetail{
		StoreId:     req.StoreId,
		ReserveId:   req.ReserveId,
		ProductId:   req.ProductId,
		QrCode:      req.QrCode,
		PayId:       req.PayId,
		CreatedBy:   req.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save buydetail")
	}
	resp := domain.BuyDetailRespone{
		StoreId:     newCust.StoreId,
		ReserveId:   newCust.ReserveId,
		ProductId:   newCust.ProductId,
		QrCode:      newCust.QrCode,
		PayId:       newCust.PayId,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s buydetailSvc) UpdateBuyDetail(id int, req domain.BuyDetailRequest) error {
	newtime := time.Now()
	cust := port.BuyDetail{

		StoreId:     req.StoreId,
		ReserveId:   req.ReserveId,
		ProductId:   req.ProductId,
		QrCode:      req.QrCode,
		PayId:       req.PayId,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update buydetail")
	}
	return nil
}
func (s buydetailSvc) DeleteBuyDetail(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete buydetail")
	}
	return nil
}
