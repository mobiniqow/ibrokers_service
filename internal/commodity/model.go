package commodity

type Commodity struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Parentid int `json:"parentId"`
    Persianname string `json:"persianName"`
    Symbol string `json:"symbol"`
}
