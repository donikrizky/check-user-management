package model

type GetRequest struct {
	Id   string `query:"id" validate:"required"`
	Name string `query:"name" validate:"required"`
}

type GetResponse struct {
	Test string `json:"test"`
}
