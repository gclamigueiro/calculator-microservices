package entity

type Request struct {
	Param1 int `json:"param1"`
	Param2 int `json:"param2"`
}

type Response struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}
