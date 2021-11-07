package core

type ResponseData struct {
	Proto      string `json:"proto"`
	Uri        string `json:"uri"`
	RealIp     string `json:"real_ip"`
	RequestURL string `json:"request_url"`
	Method     string `json:"method"`
}
