package core

type ResponseData struct {
	Proto      string `json:"proto"`
	RealIp     string `json:"real_ip"`
	Uri        string `json:"uri"`
	RequestURL string `json:"request_url"`
	Method     string `json:"method"`
	Time       string `json:"time"`
}
