package entity

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}
