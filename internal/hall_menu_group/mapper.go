package hall_menu_group




func ToHallMenuGroupResponse(buyMethod HallMenuGroup) HallMenuGroupResponse {
    return HallMenuGroupResponse {
		Id: &buyMethod.Id,
    }
}
