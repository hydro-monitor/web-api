package api_models

type NodeDTO struct {
	Id            *string `json:"id,omitempty" example:"lujan-1"`
	Description   *string `json:"description" example:"Nodo instalado debajo de un puente"`
	ManualReading *bool   `json:"manualReading" example:"false"`
	Password      *string `json:"password,omitempty" example:"aDF23DDcaF45k7J0"`
}
