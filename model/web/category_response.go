package web

/*
struct representation of response.
create response struct instead of domain (model) struct
to prevent exposing sensitive data (e.g. password),
also to follow API Specification.
*/

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
