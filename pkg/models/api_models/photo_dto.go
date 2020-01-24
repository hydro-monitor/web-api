package api_models

type PhotoDTO struct {
	ReadingId   string
	PhotoNumber int    `form:"photoNumber"`
	Photo       []byte `form:"photo"`
}
