package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/zvrvdmtr/advertising/internal/domain"
	"github.com/zvrvdmtr/advertising/internal/repository"
	"github.com/zvrvdmtr/advertising/internal/services"
)

func GetList(conn repository.DbConnection) http.HandlerFunc {
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
			pageNumber, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		ads, err := services.GetAds(conn, pageNumber)
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

func GetAd(conn repository.DbConnectionRow) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := strings.Split(r.URL.Path, "/")
		fmt.Println(path)
		elemId, _ := strconv.Atoi(path[len(path)-1])
		ad, err := services.GetAdById(conn, elemId, r.URL.Query()["fields"])
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

func CreateAd(conn repository.DbConnectionRow) http.HandlerFunc {
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

		newAd, _ := services.CreateAd(conn, ad)

		response, err := json.Marshal(newAd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}