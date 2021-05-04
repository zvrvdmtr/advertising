package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/zvrvdmtr/advertising/pkg/models"
	"github.com/zvrvdmtr/advertising/pkg/services"
)

func GetList(conn models.DbConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ads, err := services.GetAds(conn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(ads)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
		w.WriteHeader(http.StatusOK)
	}
}

func GetAd(conn models.DbConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := strings.Split(r.URL.Path, "/")
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

func CreateAd(conn models.DbConnection) http.HandlerFunc {
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

		var ad models.Ad
		err = json.Unmarshal(body, &ad)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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