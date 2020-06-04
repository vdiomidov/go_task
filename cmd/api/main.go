package main

import (
	"encoding/json"
	"io/ioutil"
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
		type Params struct {
			userId string `json:"user_id"`
		}
		var params Params

		//if !getParams(w, r, params) {
		//	return
		//}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		qqq, err := ioutil.ReadAll(r.Body)
		log.Printf(string(qqq))
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("11111111111111111")

		sessionId, err := app.GetSession(params.userId, "2222")
		if err != nil {
			log.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sessionId)
	})

	//http.HandleFunc("/price", app.GetAdvPrice)

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
