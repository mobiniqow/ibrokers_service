package hall_menu_sub_group



type CreateHallMenuSubGroupRequest struct {
	Id *int `json:"id"`
	Group *int `json:"group"`
}

type HallMenuSubGroupResponse struct {
    	Id *int `json:"id"`
	Group *int `json:"group"`
}
