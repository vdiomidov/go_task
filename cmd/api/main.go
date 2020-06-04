package main

import (
	"encoding/json"
	"log"
	"net/http"
	"task/internal/application"
	"task/internal/mysqlstorage"
)

func main() {
	storage, err := mysqlstorage.New("root:root@/adv_db")
	if err != nil {
		log.Fatal(err)
	}

	app := application.NewApp(storage)

	http.HandleFunc("/getsession", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			UserId string `json:"user_id"`
			Price  string `json:"price"`
		}

		if !getParams(w, r, &params) {
			return
		}

		sessionId, err := app.GetSession(params.UserId, params.Price)
		if err != nil {
			log.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sessionId)
	})

	http.HandleFunc("/getprice", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			SessionId string `json:"session_id"`
			Adv       []string
		}

		if !getParams(w, r, &params) {
			return
		}

		price, err := app.GetAdvPrice(params.SessionId, params.Adv)
		if err != nil {
			log.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(price)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getParams(w http.ResponseWriter, r *http.Request, params interface{}) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}
