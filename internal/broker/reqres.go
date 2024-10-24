package broker



type CreateBrokerRequest struct {
	Id *int `json:"id"`
	Spotid *int `json:"spotId"`
	Derivativesid *int `json:"derivativesId"`
}

type BrokerResponse struct {
    	Id *int `json:"id"`
	Spotid *int `json:"spotId"`
	Derivativesid *int `json:"derivativesId"`
}
