package group

type Group struct {
    Id            int    `gorm:"primary_key"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
    Parentid int `json:"parentId"`
}
