package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type userdetailSvc struct {
	repo port.UserDetailRepo
}

func NewUserDetailSvc(repo port.UserDetailRepo) domain.UserDetailSvc {
	return userdetailSvc{
		repo: repo,
	}
}

func (s userdetailSvc) GetAllUserDetail() ([]domain.UserDetailRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get userdetail form DB")
	}
	resp := []domain.UserDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.UserDetailRespone{
			UserdtId:      c.UserdtId,
			CardId:        c.CardId,
			SvcId:         c.SvcId,
			FristNameUser: c.FristNameUser,
			LastNameUser:  c.LastNameUser,
			UserAddress:   c.UserAddress,
			UserZibId:     c.UserZibId,
			UserPhone:     c.UserPhone,
			UserId:        c.UserId,
			Avatar:        c.Avatar,
			Province:      c.Province,
			CreatedBy:     c.CreatedBy,
			CreatedDate:   c.CreatedDate,
			UpdatedBy:     c.UpdatedBy,
			UpdatedDate:   c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s userdetailSvc) GetUserDetail(id int) (*domain.UserDetailRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get userdetail form DB")
	}
	resp := domain.UserDetailRespone{
		CardId:        cust.CardId,
		UserId:        cust.UserId,
		SvcId:         cust.SvcId,
		FristNameUser: cust.FristNameUser,
		LastNameUser:  cust.LastNameUser,
		UserAddress:   cust.UserAddress,
		UserZibId:     cust.UserZibId,
		UserPhone:     cust.UserPhone,
		Avatar:        cust.Avatar,
		Province:      cust.Province,
		CreatedBy:     cust.CreatedBy,
		CreatedDate:   cust.CreatedDate,
		UpdatedBy:     cust.UpdatedBy,
		UpdatedDate:   cust.UpdatedDate,
	}
	return &resp, nil
}
func (s userdetailSvc) AddUserDetail(req domain.UserDetailRequest) (*domain.UserDetailRespone, error) {
	newtime := time.Now()
	cust := port.UserDetail{
		UserId: req.UserId,
		// CardId:          req.CardId,
		// SvcId:           req.SvcId,
		// FristNameUser:   req.FristNameUser,
		// LastNameUser:    req.LastNameUser,
		// FristNameSeller: req.FristNameSeller,
		// LastNameSeller:  req.LastNameSeller,
		// UserAddress:     req.UserAddress,
		// UserZibId:       req.UserZibId,
		// TCode:           req.TCode,
		// ACode:           req.ACode,
		// ProvinceId:      req.ProvinceId,
		// UserPhone:       req.UserPhone,
		// UserEmail:       req.UserEmail,
		// UserImege:       req.UserImege,
		Avatar:      "https://static.vecteezy.com/system/resources/thumbnails/009/292/244/small/default-avatar-icon-of-social-media-user-vector.jpg",
		CreatedDate: newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save userdetail")
	}
	resp := domain.UserDetailRespone{
		CardId:        newCust.CardId,
		SvcId:         newCust.SvcId,
		FristNameUser: newCust.FristNameUser,
		LastNameUser:  newCust.LastNameUser,
		UserAddress:   newCust.UserAddress,
		UserZibId:     newCust.UserZibId,
		UserPhone:     newCust.UserPhone,
		Avatar:        newCust.Avatar,
		Province:      newCust.Province,
		UserId:        newCust.UserId,
		CreatedBy:     newCust.CreatedBy,
		CreatedDate:   newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s userdetailSvc) UpdateUserDetail(id int, req domain.UserDetailRequest) error {
	newtime := time.Now()
	cust := port.UserDetail{

		CardId: req.CardId,
		// SvcId:           req.SvcId,
		FristNameUser: req.FristNameUser,
		LastNameUser:  req.LastNameUser,
		// FristNameSeller: req.FristNameSeller,
		// LastNameSeller:  req.LastNameSeller,
		UserAddress: req.UserAddress,
		UserZibId:   req.UserZibId,
		// TCode:           req.TCode,
		// ACode:           req.ACode,
		UserPhone: req.UserPhone,
		Avatar:    req.Avatar,
		Province:  req.Province,
		// UserImege:       req.UserImege,
		// UserId:          req.UserId,
		// UpdatedBy:       req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update userdetail")
	}
	return nil
}
func (s userdetailSvc) DeleteUserDetail(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete userdetail")
	}
	return nil
}
