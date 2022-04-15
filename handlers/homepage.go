package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Homepage struct {
	l *log.Logger
}

func NewHomepage(l *log.Logger) *Homepage {
	return &Homepage{l}
}

func (h *Homepage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hit response")
}
