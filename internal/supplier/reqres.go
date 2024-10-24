package supplier



type CreateSupplierRequest struct {
	Id *int `json:"id"`
	Customerid *int `json:"customerId"`
}

type SupplierResponse struct {
    	Id *int `json:"id"`
	Customerid *int `json:"customerId"`
}
