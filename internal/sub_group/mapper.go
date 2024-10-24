package sub_group




func ToSubGroupResponse(buyMethod SubGroup) SubGroupResponse {
    return SubGroupResponse {
		Id: &buyMethod.Id,
		Parentid: &buyMethod.Parentid,
    }
}
