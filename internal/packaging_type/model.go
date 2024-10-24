package packaging_type

type PackagingType struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
}
