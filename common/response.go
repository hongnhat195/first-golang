package common

type succcessRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessReponse(data, paging, filter interface{}) *succcessRes {
	return &succcessRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *succcessRes {
	return NewSuccessReponse(data, nil, nil)
}
