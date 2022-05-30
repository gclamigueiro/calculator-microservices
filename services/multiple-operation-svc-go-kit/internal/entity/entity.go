package entity

type Request struct {
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Param1    int    `json:"param1"`
	Param2    int    `json:"param2"`
	Operation string `json:"operation"`
}

type Response struct {
	R   []int    `json:"r"`
	Err []string `json:"err,omitempty"`
}
