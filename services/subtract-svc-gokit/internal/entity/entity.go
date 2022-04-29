package entity

type SubtractRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SubtractResponse struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}
