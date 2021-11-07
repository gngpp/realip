package core

import (
	"net"
	"net/http"
	"strings"
)

func ExtractRealIp(request *http.Request) string {
	if strings.Contains(request.RemoteAddr, "::1") || strings.Contains(request.RemoteAddr, "localhost") {
		return "127.0.0.1"
	}
	// Obtain through firewall
	ip := strings.TrimSpace(request.Header.Get("X-Forwarded-For"))
	if ip != "" {
		index := strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = strings.TrimSpace(ip[:index]); ip != "" {
			return ip
		}
	}
	if ip = strings.TrimSpace(request.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}
	ip, _, _ = net.SplitHostPort(request.RemoteAddr)
	return ip
}

func GetRequestURL(request *http.Request) (url string) {
	if request.TLS == nil {
		//goland:noinspection HttpUrlsUsage
		url = "http://" + request.Host + request.RequestURI
	} else {
		url = "https://" + request.Host + request.RequestURI
	}
	return
}
