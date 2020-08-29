package api_models

type PaginatedDTO struct {
	PageState []byte
	Elements  []*GetReadingDTO
}
