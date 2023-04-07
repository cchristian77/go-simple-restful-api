package web

// standardized response

type WebResponse struct {
	Code   int
	Status string
	Data   interface{} // to handle any data type
}
