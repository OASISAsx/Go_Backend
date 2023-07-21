package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type cartSvc struct {
	repo port.CartRepo
}

func NewCartSvc(repo port.CartRepo) domain.CartSvc {
	return cartSvc{
		repo: repo,
	}
}

func (s cartSvc) GetAllCart() ([]domain.CartRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Cart form DB")
	}
	resp := []domain.CartRespone{}
	for _, c := range custs {
		resp = append(resp, domain.CartRespone{
			CartId:     c.CartId,
			SvcId:         c.SvcId,
			ProductName:   c.ProductName,
			ProductDesc:   c.ProductDesc,
			ProductImages: c.ProductImages,
			Producttype: c.Producttype,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s cartSvc) GetById(id int) (*domain.CartRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Cart form DB")
	}
	resp := domain.CartRespone{
		CartId:     cust.CartId,
		SvcId:         cust.SvcId,
		ProductName:   cust.ProductName,
		ProductDesc:   cust.ProductDesc,
		ProductImages: cust.ProductImages,
		Producttype:   cust.Producttype,
		CreatedBy:     cust.CreatedBy,
		CreatedDate:   cust.CreatedDate,
		UpdatedBy:     cust.UpdatedBy,
		UpdatedDate:   cust.UpdatedDate,
	}
	return &resp, nil
}
func (s cartSvc) AddCart(req domain.CartRequest) (*domain.CartRespone, error) {
	newtime := time.Now()
	cust := port.Cart{
		SvcId:         req.SvcId,
		ProductName:   req.ProductName,
		ProductDesc:   req.ProductDesc,
		ProductImages: req.ProductImages,
		Producttype:  req.Producttype,
		CreatedBy:     req.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Cart")
	}
	resp := domain.CartRespone{
		SvcId:         newCust.SvcId,
		ProductName:   newCust.ProductName,
		ProductDesc:   newCust.ProductDesc,
		ProductImages: newCust.ProductImages,
		Producttype:   newCust.Producttype,
		CreatedBy:     newCust.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s cartSvc) UpdateCart(id int, req domain.CartRequest) error {
	newtime := time.Now()
	cust := port.Cart{
		SvcId:         req.SvcId,
		ProductName:   req.ProductName,
		ProductDesc:   req.ProductDesc,
		ProductImages: req.ProductImages,
		Producttype:   req.Producttype,
		UpdatedBy:     req.UpdatedBy,
		UpdatedDate:   newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Cart")
	}
	return nil
}
func (s cartSvc) DeleteCart(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Cart")
	}
	return nil
}
