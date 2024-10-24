package group




func ToGroupResponse(buyMethod Group) GroupResponse {
    return GroupResponse {
		Id: &buyMethod.Id,
		Parentid: &buyMethod.Parentid,
    }
}
