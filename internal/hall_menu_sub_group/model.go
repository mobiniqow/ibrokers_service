package hall_menu_sub_group

type HallMenuSubGroup struct {
    Id            int    `gorm:"primary_key"`
    Name string `json:"name"`
    Group int `json:"group"`
}
