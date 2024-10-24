package broker

type Broker struct {
    Id            int    `form:"id" gorm:"primary_key"`
    Description string `form:"description"`
    Persianname string `form:"persianName"`
    Spotid int `form:"spotId"`
    Derivativesid int `form:"derivativesId"`
    Nationalid string `form:"nationalId"`
}
