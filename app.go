package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"realip/core"
)

//goland:noinspection SpellCheckingInspection
func main() {
	port := port()
	http.HandleFunc("/ip", requestHandler)
	err := http.ListenAndServe(port, http.DefaultServeMux)
	if err != nil {
		log.Fatalf("abnormal listening address, %s", err)
	}
}

//goland:noinspection SpellCheckingInspection
func port() string {
	port := flag.String("p", "8080", "realip server port")
	// parse variable
	flag.Parse()
	log.Println("realip server port:" + *port)
	return ":" + *port
}

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	handlerTimeFunc := core.GetHandlerTime()
	url := core.GetRequestURL(request)
	ip := core.ExtractRealIp(request)
	data := &core.ResponseData{
		Proto:      request.Proto,
		RealIp:     ip,
		Uri:        request.RequestURI,
		RequestURL: url,
		Method:     request.Method,
	}
	defer func() {
		data.Time = handlerTimeFunc()
		marshal, err := json.Marshal(data)
		writer.Header().Add("Content-type", "application/json")
		if err != nil {
			writer.WriteHeader(500)
		}
		_, err = fmt.Fprint(writer, string(marshal))
		if err != nil {
			log.Println(err)
		}
	}()
}
