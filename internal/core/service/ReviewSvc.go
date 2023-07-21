package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
	"time"
)

type reviewSvc struct {
	repo port.ReviewRepo
}

func NewReviewSvc(repo port.ReviewRepo) domain.ReviewSvc {
	return reviewSvc{
		repo: repo,
	}
}

func (s reviewSvc) GetAllReview() ([]domain.ReviewRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get review form DB")
	}
	resp := []domain.ReviewRespone{}
	for _, c := range custs {
		resp = append(resp, domain.ReviewRespone{
			RvId:        c.RvId,
			UserId:      c.UserId,
			ProductId:   c.ProductId,
			RvImg:       c.RvImg,
			RvRank:      c.RvRank,
			RvComment:   c.RvComment,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s reviewSvc) GetReview(id int) (*domain.ReviewRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get review form DB")
	}
	resp := domain.ReviewRespone{
		RvId:        cust.RvId,
		UserId:      cust.UserId,
		ProductId:   cust.ProductId,
		RvImg:       cust.RvImg,
		RvRank:      cust.RvRank,
		RvComment:   cust.RvComment,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}
func (s reviewSvc) AddReview(req domain.ReviewRequest) (*domain.ReviewRespone, error) {
	newtime := time.Now()
	cust := port.Review{
		UserId:      req.UserId,
		ProductId:   req.ProductId,
		RvImg:       req.RvImg,
		RvRank:      req.RvRank,
		RvComment:   req.RvComment,
		CreatedBy:   req.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save review")
	}
	resp := domain.ReviewRespone{
		UserId:      newCust.UserId,
		ProductId:   newCust.ProductId,
		RvImg:       newCust.RvImg,
		RvRank:      newCust.RvRank,
		RvComment:   newCust.RvComment,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: newtime.Format(time.DateTime),
	}

	return &resp, nil
}
func (s reviewSvc) UpdateReview(id int, req domain.ReviewRequest) error {
	newtime := time.Now()
	cust := port.Review{
		UserId:      req.UserId,
		ProductId:   req.ProductId,
		RvImg:       req.RvImg,
		RvRank:      req.RvRank,
		RvComment:   req.RvComment,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: newtime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update review")
	}
	return nil
}
func (s reviewSvc) DeleteReview(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete review")
	}
	return nil
}
