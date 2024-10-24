package group_hall



type CreateGroupHallRequest struct {
	Id *int `json:"id"`
	Group *int `json:"group"`
	Hall *int `json:"hall"`
}

type GroupHallResponse struct {
    	Id *int `json:"id"`
	Group *int `json:"group"`
	Hall *int `json:"hall"`
}
