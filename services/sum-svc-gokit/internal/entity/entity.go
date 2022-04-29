package entity

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}
