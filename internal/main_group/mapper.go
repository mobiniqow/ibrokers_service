package main_group




func ToMainGroupResponse(buyMethod MainGroup) MainGroupResponse {
    return MainGroupResponse {
		Id: &buyMethod.Id,
    }
}
