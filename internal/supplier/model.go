package supplier

type Supplier struct {
    Id            int    `gorm:"primary_key"`
    Customerid int `json:"customerId"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
    Nationalcode string `json:"nationalCode"`
}
