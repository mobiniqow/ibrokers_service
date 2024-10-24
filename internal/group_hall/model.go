package group_hall

type GroupHall struct {
    Id            int    `gorm:"primary_key"`
    Group int `json:"group"`
    Hall int `json:"hall"`
}
