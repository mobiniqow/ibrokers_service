package sub_group

type SubGroup struct {
    Id            int    `gorm:"primary_key"`
    Parentid int `json:"parentId"`
    Description string `json:"description"`
    Persianname string `json:"persianName"`
}
