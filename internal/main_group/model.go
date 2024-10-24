package main_group

type MainGroup struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
    Icon string `json:"icon"`
}
