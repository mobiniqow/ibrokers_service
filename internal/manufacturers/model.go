package manufacturers

type Manufacturers struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Symbol string `json:"symbol"`
    Persianname string `json:"persianName"`
}
