package domain

type ReviewSvc interface {
	GetAllReview() ([]ReviewRespone, error)
	GetReview(int) (*ReviewRespone, error)
	AddReview(ReviewRequest) (*ReviewRespone, error)
	UpdateReview(int, ReviewRequest) error
	DeleteReview(int) error
}

type ReviewRequest struct {
	
	ProductId   uint   `json:"productid"`
	UserId      uint   `json:"userid"`
	RvId        uint   `json:"rvid"`
	RvImg       string `json:"rvimg"`
	RvRank      string `json:"rvrank"`
	RvComment   string `json:"rvcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type ReviewRespone struct {
	ProductId   uint   `json:"productid"`
	RvId        uint   `json:"rvid"`
	UserId      uint   `json:"userid"`
	RvImg       string `json:"rvimg"`
	RvRank      string `json:"rvrank"`
	RvComment   string `json:"rvcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
