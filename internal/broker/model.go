package broker

type Broker struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
    Spotid int `json:"spotId"`
    Derivativesid int `json:"derivativesId"`
    Nationalid string `json:"nationalId"`
}
