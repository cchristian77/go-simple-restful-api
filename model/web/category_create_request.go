package web

// struct representation of request

type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,min=1" json:"name"`
}
