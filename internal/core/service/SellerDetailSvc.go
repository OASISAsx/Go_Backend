package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type sellerDetailSvc struct {
	repo port.SellerDetailRepo
}

func NewSellerDetailSvc(repo port.SellerDetailRepo) domain.SellerDetailSvc {
	return sellerDetailSvc{
		repo: repo,
	}
}

func (s sellerDetailSvc) GetAllSellerDetail() ([]domain.SellerDetailRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get sellers detail form DB")
	}
	resp := []domain.SellerDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.SellerDetailRespone{
			UserdeId:     c.UserdeId,
			UserId:       c.UserId,
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Phone:        c.Phone,
			BankName:     c.BankName,
			BankId:       c.BankId,
			PersonCard:   c.PersonCard,
			RecordStatus: c.RecordStatus,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s sellerDetailSvc) GetSellerDetail(id int) (*domain.SellerDetailRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get SellerDetail form DB")
	}
	resp := domain.SellerDetailRespone{
		UserdeId:     cust.UserdeId,
		UserId:       cust.UserId,
		FirstName:    cust.FirstName,
		LastName:     cust.LastName,
		Phone:        cust.Phone,
		BankName:     cust.BankName,
		BankId:       cust.BankId,
		PersonCard:   cust.PersonCard,
		RecordStatus: cust.RecordStatus,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}

func (r sellerDetailSvc) AddSellerDetail(req domain.SellerDetailRequest) (*domain.SellerDetailRespone, error) {
	currentTime := time.Now()
	cust := port.SellerDetail{
		UserId:       req.UserId,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		BankName:     req.BankName,
		BankId:       req.BankId,
		RecordStatus: "รออนุมัติ",
		PersonCard:   req.PersonCard,
		CreatedDate:  currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save SellerDetail	")
	}
	resp := domain.SellerDetailRespone{
		UserId:      newCust.UserId,
		FirstName:   newCust.FirstName,
		LastName:    newCust.LastName,
		Phone:       newCust.Phone,
		BankName:    newCust.BankName,
		BankId:      newCust.BankId,
		PersonCard:  newCust.PersonCard,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s sellerDetailSvc) UpdateSellerDetail(id int, req domain.SellerDetailRequest) error {
	currentTime := time.Now()
	cust := port.SellerDetail{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		BankName:  req.BankName,
		BankId:    req.BankId,
		// PersonCard:   req.PersonCard,
		RecordStatus: req.RecordStatus,
		UpdatedBy:    req.UpdatedBy,
		UpdatedDate:  currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update SellerDetail: ")
	}
	return nil
}
func (s sellerDetailSvc) DeleteSellerDetail(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete SellerDetail")
	}
	return nil
}
