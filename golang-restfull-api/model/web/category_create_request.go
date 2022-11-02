package web

type CategoryCreatRequest struct {
	Name string `json:"name" validate:"required,max=200,min=1"`
}
