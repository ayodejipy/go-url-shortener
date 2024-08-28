package handler

import (
	"fmt"
	"net/http"
	"rest/api/internals/service"
)

type pingHandler struct {
	service *service.PingService
}

func (h *pingHandler) sayHello(w http.ResponseWriter, r *http.Request) {
	val := h.service.Ping()
	fmt.Printf("%v\n", val)
}