package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"strings"

	"github.com/zvrvdmtr/advertising/internal/domain"
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

func GetList(service domain.AdServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// FIXME: Bad codestyle
		page := r.URL.Query().Get("page")
		var pageNumber int
		var err error
		if page != "" {
			pageNumber, err = strconv.Atoi(page)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		ads, err := service.GetAds(pageNumber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(ads)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func GetAd(service domain.AdServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := strings.Split(r.URL.Path, "/")
		fmt.Println(path)
		elemId, _ := strconv.Atoi(path[len(path)-1])
		ad, err := service.GetAdById(elemId, r.URL.Query()["fields"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response, err :=  json.Marshal(ad)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func CreateAd(service domain.AdServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var ad domain.Ad
		err = json.Unmarshal(body, &ad)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		newAd, err := service.CreateAd(ad)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(newAd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func NewAdHandler(service domain.AdServiceInterface) *Handler{
	handler := Handler{DefaultRouter: func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}}
	handler.HandleFunc("^/ad/[0-9]+$", GetAd(service))
	handler.HandleFunc("/ads", GetList(service))
	handler.HandleFunc("/create", CreateAd(service))
	return &handler
}