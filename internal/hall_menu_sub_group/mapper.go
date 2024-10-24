package hall_menu_sub_group




func ToHallMenuSubGroupResponse(buyMethod HallMenuSubGroup) HallMenuSubGroupResponse {
    return HallMenuSubGroupResponse {
		Id: &buyMethod.Id,
		Group: &buyMethod.Group,
    }
}
