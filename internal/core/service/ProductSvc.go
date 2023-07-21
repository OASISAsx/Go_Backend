package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
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
			Producttype:   c.Producttype,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
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
		Producttype:   cust.Producttype,
		CreatedBy:     cust.CreatedBy,
		CreatedDate:   cust.CreatedDate,
		UpdatedBy:     cust.UpdatedBy,
		UpdatedDate:   cust.UpdatedDate,
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
		ProductStock:  req.ProductStock,
		ProductImages: req.ProductImages,
		Producttype:   req.Producttype,
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
		Producttype:   newCust.Producttype,
		CreatedBy:     newCust.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
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
		Producttype:   req.Producttype,
		UpdatedBy:     req.UpdatedBy,
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
			Producttype:   c.Producttype,
			ProductImages: c.ProductImages,
			ProductPrice:  c.ProductPrice,
			ProductDesc:   c.ProductDesc,
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

func (s productSvc) UpdateStatusProduct(id int, req domain.StatusProduct) error {
	cust := port.Product{
		Status: req.Status,
	}
	err := s.repo.UpdateStatusProduct(id, cust.Status)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Status: ")
	}
	return nil
}
func (s productSvc) UpdateCount(id int, req domain.ProductRequest) error {
	currentTime := time.Now()
	cust := port.Product{
		Count:       +1,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.UpdateCount(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Count: ")
	}
	return nil
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
