package domain

type ReviewSvc interface {
	GetAllReview() ([]ReviewRespone, error)
	GetReview(int) (*ReviewRespone, error)
	AddReview(ReviewRequest) (*ReviewRespone, error)
	UpdateReview(int, ReviewRequest) error
	DeleteReview(int) error
}

type ReviewRequest struct {
	UserId      uint   `json:"userid"`
	BuyerId     uint   `json:"buyerid"`
	RvRank      string `json:"rvrank"`
	RvComment   string `json:"rvcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type ReviewRespone struct {
	RvId        uint   `json:"rvid"`
	UserId      uint   `json:"userid"`
	BuyerId     uint   `json:"buyerid"`
	RvRank      string `json:"rvrank"`
	RvComment   string `json:"rvcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}