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
			PayId:         c.PayId,
			UserId:        c.UserId,
			ProductId:     c.ProductId,
			ById:          c.ById,
			PaySlip:       c.PaySlip,
			ParNum:        c.ParNum,
			TransPort:     c.TransPort,
			PayStatus:     c.PayStatus,
			ProductName:   c.ProductName,
			ProductImages: c.ProductImages,
			ProductPrice:  c.ProductPrice,
			ProductStock:  c.ProductStock,
			ProductType:   c.ProductType,
			ProductDesc:   c.ProductDesc,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return resp, nil
}
func (s buydetailSvc) SearchBuyDetail(name string) (*[]domain.BuyDetailRespone, error) {
	custs, err := s.repo.Search(name)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get payment form DB")
	}
	resp := []domain.BuyDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BuyDetailRespone{
			PayId:         c.PayId,
			UserId:        c.UserId,
			ProductId:     c.ProductId,
			PaySlip:       c.PaySlip,
			ParNum:        c.ParNum,
			TransPort:     c.TransPort,
			PayStatus:     c.PayStatus,
			ProductName:   c.ProductName,
			ProductImages: c.ProductImages,
			ProductPrice:  c.ProductPrice,
			ProductStock:  c.ProductStock,
			ProductType:   c.ProductType,
			ProductDesc:   c.ProductDesc,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})
	}
	return &resp, nil
}

func (s buydetailSvc) GetBuyDetail(id int) (*domain.BuyDetailRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get buydetail form DB")
	}
	resp := domain.BuyDetailRespone{

		PayId:         cust.PayId,
		UserId:        cust.UserId,
		ById:          cust.ById,
		ProductId:     cust.ProductId,
		PaySlip:       cust.PaySlip,
		PayStatus:     cust.PayStatus,
		ParNum:        cust.ParNum,
		TransPort:     cust.TransPort,
		ProductName:   cust.ProductName,
		ProductImages: cust.ProductImages,
		ProductPrice:  cust.ProductPrice,
		ProductStock:  cust.ProductStock,
		ProductType:   cust.ProductType,
		ProductDesc:   cust.ProductDesc,
		CreatedBy:     cust.CreatedBy,
		CreatedDate:   cust.CreatedDate,
		UpdatedBy:     cust.UpdatedBy,
		UpdatedDate:   cust.UpdatedDate,
	}
	return &resp, nil
}
func (s buydetailSvc) AddBuyDetail(req domain.BuyDetailRequest) (*domain.BuyDetailRespone, error) {
	newtime := time.Now()
	cust := port.BuyDetail{
		UserId:        req.UserId,
		ProductId:     req.ProductId,
		ById:          req.ById,
		PaySlip:       req.PaySlip,
		PayStatus:     "กำลังดำเนินการ",
		ParNum:        req.ParNum,
		TransPort:     req.TransPort,
		ProductName:   req.ProductName,
		ProductImages: req.ProductImages,
		ProductPrice:  req.ProductPrice,
		ProductStock:  req.ProductStock,
		ProductType:   req.ProductType,
		ProductDesc:   req.ProductDesc,
		CreatedBy:     req.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save buydetail")
	}
	resp := domain.BuyDetailRespone{
		UserId:        newCust.UserId,
		ProductId:     newCust.ProductId,
		PaySlip:       newCust.PaySlip,
		ParNum:        newCust.ParNum,
		TransPort:     newCust.TransPort,
		PayStatus:     newCust.PayStatus,
		ProductName:   newCust.ProductName,
		ProductImages: newCust.ProductImages,
		ProductPrice:  newCust.ProductPrice,
		ProductStock:  newCust.ProductStock,
		ProductType:   newCust.ProductType,
		ProductDesc:   newCust.ProductDesc,
		CreatedBy:     newCust.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s buydetailSvc) UpdateBuyDetail(id int, req domain.BuyDetailRequest) error {
	newtime := time.Now()
	cust := port.BuyDetail{

		ParNum:       req.ParNum,
		TransPort:    req.TransPort,
		PayStatus:    req.PayStatus,
		ProductStock: req.ProductStock,
		UpdatedBy:    "Admin",
		UpdatedDate:  newtime.Format(time.DateTime),
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
func (s buydetailSvc) GetAllBuyDetailId(id int) ([]domain.BuyDetailRespone, error) {
	custs, err := s.repo.GetAllId(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get payment form DB")
	}
	resp := []domain.BuyDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BuyDetailRespone{
			PayId:         c.PayId,
			UserId:        c.UserId,
			ProductId:     c.ProductId,
			ById:          c.ById,
			PaySlip:       c.PaySlip,
			ParNum:        c.ParNum,
			TransPort:     c.TransPort,
			PayStatus:     c.PayStatus,
			ProductName:   c.ProductName,
			ProductImages: c.ProductImages,
			ProductPrice:  c.ProductPrice,
			ProductStock:  c.ProductStock,
			ProductType:   c.ProductType,
			ProductDesc:   c.ProductDesc,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return resp, nil
}
func (s buydetailSvc) GetAllUserId(id int) ([]domain.BuyDetailRespone, error) {
	custs, err := s.repo.GetAllUserId(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get payment form DB")
	}
	resp := []domain.BuyDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BuyDetailRespone{
			PayId:         c.PayId,
			UserId:        c.UserId,
			ProductId:     c.ProductId,
			ById:          c.ById,
			PaySlip:       c.PaySlip,
			ParNum:        c.ParNum,
			TransPort:     c.TransPort,
			PayStatus:     c.PayStatus,
			ProductName:   c.ProductName,
			ProductImages: c.ProductImages,
			ProductPrice:  c.ProductPrice,
			ProductType:   c.ProductType,
			ProductStock:  c.ProductStock,
			ProductDesc:   c.ProductDesc,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return resp, nil
}
