package service

import (
	"collection/errs"
	"collection/internal/core/domain"
	"collection/internal/core/port"
	"net/http"
)

type reportSvc struct {
	repo port.ReportRepo
}

func NewReportSvc(repo port.ReportRepo) domain.ReportSvc {
	return reportSvc{
		repo: repo,
	}
}

func (s reportSvc) GetAllReport() ([]domain.ReportResp, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get report form DB")
	}
	resp := []domain.ReportResp{}
	for _, c := range custs {
		resp = append(resp, domain.ReportResp{
			ReportId: c.ReportId,
			Username: c.Username,
			Email:    c.Email,
			Message:  c.Message,
		})

	}
	return resp, nil
}

func (s reportSvc) GetReport(id int) (*domain.ReportResp, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get report form DB")
	}
	resp := domain.ReportResp{
		ReportId: cust.ReportId,
	}
	return &resp, nil
}
func (s reportSvc) AddReport(req domain.ReportReq) (*domain.ReportResp, error) {
	cust := port.Report{
		Username: req.Username,
		Email:    req.Email,
		Message:  req.Message,
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save report")
	}
	resp := domain.ReportResp{
		ReportId: newCust.ReportId,
		Username: newCust.Username,
		Email:    newCust.Email,
		Message:  newCust.Message,
	}

	return &resp, nil
}
func (s reportSvc) UpdateReport(id int, req domain.ReportReq) error {
	cust := port.Report{
		Username: req.Username,
		Email:    req.Email,
		Message:  req.Message,
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update report")
	}
	return nil
}
func (s reportSvc) DeleteReport(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete report")
	}
	return nil
}
