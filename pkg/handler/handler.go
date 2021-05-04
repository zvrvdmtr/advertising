package handler

import (
	"net/http"
	"regexp"
)


type Route struct {
	Pattern *regexp.Regexp
	Handler func(w http.ResponseWriter, r *http.Request)
}

type Handler struct{
	Routes []Route
	DefaultRouter func(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.Routes {
		if matches := route.Pattern.FindStringSubmatch(r.URL.Path); len(matches) > 0 {
			route.Handler(w, r)
			return
		}
	}
	h.DefaultRouter(w, r)
}

func (h *Handler) HandleFunc(pattern string, hander func(w http.ResponseWriter, r *http.Request)) {
	regexpPattern := regexp.MustCompile(pattern)
	h.Routes = append(h.Routes, Route{regexpPattern, hander})
}

func NewHandler() *Handler{
	return &Handler{DefaultRouter: func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}}
}