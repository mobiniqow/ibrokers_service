package sub_group



type CreateSubGroupRequest struct {
	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}

type SubGroupResponse struct {
    	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}
