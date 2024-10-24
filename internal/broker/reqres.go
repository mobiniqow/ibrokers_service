package broker



type CreateBrokerRequest struct {
	Id *int `form:"id"`
	Description *string `form:"description"`
	Persianname *string `form:"persianName"`
	Spotid *int `form:"spotId"`
	Derivativesid *int `form:"derivativesId"`
	Nationalid *string `form:"nationalId"`
}

type BrokerResponse struct {
    	Id *int `form:"id"`
	Description *string `form:"description"`
	Persianname *string `form:"persianName"`
	Spotid *int `form:"spotId"`
	Derivativesid *int `form:"derivativesId"`
	Nationalid *string `form:"nationalId"`
}
