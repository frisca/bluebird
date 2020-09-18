package dbmodels

type OrderReq struct {
	Topic string      `json:"topic"`
	Data  interface{} `json:"data"`
}