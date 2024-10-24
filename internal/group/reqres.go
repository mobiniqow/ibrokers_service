package group



type CreateGroupRequest struct {
	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}

type GroupResponse struct {
    	Id *int `json:"id"`
	Parentid *int `json:"parentId"`
}
