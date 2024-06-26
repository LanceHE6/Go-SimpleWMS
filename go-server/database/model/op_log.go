package model

type OPLog struct {
	Base
	Uid       string `json:"uid"`
	Resource  string `json:"resource"`
	Operation string `json:"operation"`
	Status    string `json:"status"`
}
