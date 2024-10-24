package hall_menu_group

type HallMenuGroup struct {
    Id            int    `gorm:"primary_key"`
    Name string `json:"name"`
}
