package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"fmt"
	"net/http"
	"time"
)

type productSvc struct {
	repo port.ProductRepo
}

func NewProductSvc(repo port.ProductRepo) domain.ProductSvc {
	return productSvc{
		repo: repo,
	}
}

func (s productSvc) GetAllProduct() ([]domain.ProductRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get product form DB")
	}
	resp := []domain.ProductRespone{}
	for _, c := range custs {
		resp = append(resp, domain.ProductRespone{
			ProductId:     c.ProductId,
			SvcId:         c.SvcId,
			ProductName:   c.ProductName,
			ProductDesc:   c.ProductDesc,
			ProductPrice:  c.ProductPrice,
			ProductStock:  c.ProductStock,
			ProductImages: c.ProductImages,
			ProductImagex: c.ProductImagex,
			ProductImagey: c.ProductImagey,
			ProductImagez: c.ProductImagez,
			ProductType:   c.ProductType,
			RvRank:        c.RvRank,
			Status:        c.Status,
			SellStatus:    c.SellStatus,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
			ProductQr:     fmt.Sprintf("https://upload.wikimedia.org/wikipedia/commons/d/d0/QR_code_for_mobile_English_Wikipedia.svg"),
		})

	}
	return resp, nil
}

func (s productSvc) GetById(id int) (*domain.ProductRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get product form DB")
	}
	resp := domain.ProductRespone{
		ProductId:     cust.ProductId,
		SvcId:         cust.SvcId,
		ProductName:   cust.ProductName,
		ProductPrice:  cust.ProductPrice,
		ProductStock:  cust.ProductStock,
		ProductDesc:   cust.ProductDesc,
		ProductImages: cust.ProductImages,
		ProductImagex: cust.ProductImagex,
		ProductImagey: cust.ProductImagey,
		ProductImagez: cust.ProductImagez,
		ProductType:   cust.ProductType,
		RvRank:        cust.RvRank,
		Status:        cust.Status,
		CreatedBy:     cust.CreatedBy,
		CreatedDate:   cust.CreatedDate,
		UpdatedBy:     cust.UpdatedBy,
		UpdatedDate:   cust.UpdatedDate,
		ProductQr:     fmt.Sprintf("https://upload.wikimedia.org/wikipedia/commons/d/d0/QR_code_for_mobile_English_Wikipedia.svg"),
	}
	return &resp, nil
}
func (s productSvc) AddProduct(req domain.ProductRequest) (*domain.ProductRespone, error) {
	newtime := time.Now()
	cust := port.Product{
		SvcId:         req.SvcId,
		ProductName:   req.ProductName,
		ProductDesc:   req.ProductDesc,
		ProductPrice:  req.ProductPrice,
		ProductStock:  "1",
		ProductImages: req.ProductImages,
		ProductImagex: req.ProductImagex,
		ProductImagey: req.ProductImagey,
		ProductImagez: req.ProductImagez,
		RvRank:        "0",
		Status:        "กำลังดำเนินการ",
		SellStatus:    false,
		ProductType:   req.ProductType,
		CreatedBy:     req.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save product")
	}
	resp := domain.ProductRespone{
		SvcId:         newCust.SvcId,
		ProductName:   newCust.ProductName,
		ProductDesc:   newCust.ProductDesc,
		ProductPrice:  newCust.ProductPrice,
		ProductStock:  newCust.ProductStock,
		ProductImages: newCust.ProductImages,
		ProductImagex: newCust.ProductImagex,
		ProductImagey: newCust.ProductImagey,
		ProductImagez: newCust.ProductImagez,
		RvRank:        newCust.RvRank,
		Status:        newCust.Status,
		ProductType:   newCust.ProductType,
		CreatedBy:     newCust.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
		ProductQr:     fmt.Sprintf("https://upload.wikimedia.org/wikipedia/commons/d/d0/QR_code_for_mobile_English_Wikipedia.svg"),
	}

	return &resp, nil
}
func (s productSvc) UpdateProduct(id int, req domain.ProductRequest) error {
	newtime := time.Now()
	cust := port.Product{
		SvcId:         req.SvcId,
		ProductName:   req.ProductName,
		ProductDesc:   req.ProductDesc,
		ProductPrice:  req.ProductPrice,
		ProductStock:  req.ProductStock,
		ProductImages: req.ProductImages,
		ProductImagex: req.ProductImagex,
		ProductImagey: req.ProductImagey,
		ProductImagez: req.ProductImagez,
		ProductType:   req.ProductType,
		RvRank:        req.RvRank,
		Status:        req.Status,
		UpdatedBy:     "Admin",
		UpdatedDate:   newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update product")
	}
	return nil
}
func (s productSvc) DeleteProduct(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete product")
	}
	return nil
}
func (s productSvc) Search(productName string) (*[]domain.ProductRespone, error) {
	custs, err := s.repo.Search(productName)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get product form DB")
	}
	resp := []domain.ProductRespone{}
	for _, c := range custs {
		resp = append(resp, domain.ProductRespone{
			ProductId:     c.ProductId,
			ProductName:   c.ProductName,
			ProductType:   c.ProductType,
			ProductImages: c.ProductImages,
			ProductImagex: c.ProductImagex,
			ProductImagey: c.ProductImagey,
			ProductImagez: c.ProductImagez,
			ProductPrice:  c.ProductPrice,
			ProductDesc:   c.ProductDesc,
			RvRank:        c.RvRank,
			Status:        c.Status,
			SellStatus:    c.SellStatus,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return &resp, nil
}

func (s productSvc) UpdateSellStatus(id int, req domain.SellStatusProduct) error {
	cust := port.Product{
		SellStatus: req.SellStatus,
	}
	err := s.repo.UpdateSellStatus(id, cust.SellStatus)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Status: ")
	}
	return nil
}
