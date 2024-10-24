package offer_type

type OfferType struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
}
