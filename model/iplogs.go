package model

// Ip represents data about a log record.
type Ip struct {
	IP     string  `json:"IP"` // json 回傳時IP 的 Key 對應到 IP
	Time   string  `json:"時間"` // json 回傳時Time 的 Key 對應到 時間
	Url    string  `json:"網址"`
	Status float64 `json:"狀態"`
}
