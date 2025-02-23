package models

type Req struct{
	ID uint64 `json:"id"`
}

type Resp struct{
	Error string `json:"error"`
	Status int `json:"status"`
}

func NewResp(err error, code int) *Resp {
	return &Resp{
		Error: err.Error(),
		Status: code,
	}
}