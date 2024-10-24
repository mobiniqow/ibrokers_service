package buy_method



type CreateBuyMethodRequest struct {
	Id *int `form:"id"`
	Description *string `form:"description"`
	Persianname *string `form:"persianName"`
}

type BuyMethodResponse struct {
    	Id *int `form:"id"`
	Description *string `form:"description"`
	Persianname *string `form:"persianName"`
}
