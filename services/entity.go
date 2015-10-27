package services

type Entity struct {
	Id   uint64      `json:"id"`
	Href string      `json:"href"`
	Data interface{} `json:"data"`
}
