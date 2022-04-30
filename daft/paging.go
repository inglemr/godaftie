package daft

type Paging struct {
	TotalPages     int `json:"totalPages"`
	CurrentPage    int `json:"currentPage"`
	NextFrom       int `json:"nextFrom"`
	PreviousFrom   int `json:"previousFrom"`
	DisplayingFrom int `json:"displayingFrom"`
	DisplayingTo   int `json:"displayingTo"`
	TotalResults   int `json:"totalResults"`
	PageSize       int `json:"pageSize"`
}
