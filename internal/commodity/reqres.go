package commodity



type CreateCommodityRequest struct {
	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}

type CommodityResponse struct {
    	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}
