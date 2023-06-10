package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type httpResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func wrap(app *App, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "app", app)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func serveResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	out, err := json.Marshal(httpResp{Status: "success", Data: data})
	if err != nil {
		sendErrorResponse(w, "Internal Server Error", http.StatusInternalServerError, nil)
	}

	w.Write(out)
}

func sendErrorResponse(w http.ResponseWriter, message string, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)

	out, _ := json.Marshal(httpResp{Status: "error", Message: message, Data: data})

	w.Write(out)
}
