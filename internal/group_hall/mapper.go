package group_hall




func ToGroupHallResponse(buyMethod GroupHall) GroupHallResponse {
    return GroupHallResponse {
		Id: &buyMethod.Id,
		Group: &buyMethod.Group,
		Hall: &buyMethod.Hall,
    }
}
