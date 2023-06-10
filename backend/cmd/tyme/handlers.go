package main

import (
	"net/http"
)

func testRoute(w http.ResponseWriter, r *http.Request) {
	serveResponse(w, map[string]interface{}{"message": "Server Up and Running"})
}
