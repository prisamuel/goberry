package main

import (
	"net/http"
	"time"
)

func logRequest(r *http.Request) {
	logger.Log("channel", "request", "service", serviceName, "method", r.Method, "url", r.URL.String(), "headers", r.Header, "ts", time.Now())
}

func logChannel(channel, message string) {
	logger.Log("channel", channel, "service", serviceName, "message", message, "ts", time.Now())
}
