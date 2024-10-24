package buy_method

type BuyMethod struct {
    Id            int    `form:"id" gorm:"primary_key"`
    Description string `form:"description"`
    Persianname string `form:"persianName"`
}
