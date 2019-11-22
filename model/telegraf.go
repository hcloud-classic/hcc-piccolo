package model

// Telegraf - cgs
type Telegraf struct {
	Id   string   `json:"id"`
	Data []Series `json:"data"`
}

// Series - cgs
type Series struct {
	Time  int `json:"x"`
	Value int `json:"y"`
}
