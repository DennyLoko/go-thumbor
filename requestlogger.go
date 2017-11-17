package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type RequestLogger struct {
	Handler http.Handler
}

func (rl *RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.Handler.ServeHTTP(w, r)

	var status, lenght int = 404, 0

	if etag := w.Header().Get("Etag"); etag != "" {
		status = 304

		if l := w.Header().Get("Content-Length"); l != "" {
			status = 200
			lenght, _ = strconv.Atoi(l)
		}
	}

	ip, _ := rl.userIP(r)

	fmt.Println(fmt.Sprintf(
		"%s - - [%s] \"%s /%s %s\" %d %d \"-\" \"%s\"",
		ip,
		time.Now().Format("2/Jan/2006:15:04:05 -0700"),
		r.Method,
		r.URL.String(),
		r.Proto,
		status,
		lenght,
		r.UserAgent()))
}

func (*RequestLogger) userIP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	return ip, nil

}
