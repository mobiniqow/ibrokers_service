package currency_unit

type CurrencyUnit struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
}
