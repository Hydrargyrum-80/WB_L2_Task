package serverControl

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

func (c *Controller) CheckDate(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		_, err := time.Parse(time.DateOnly, date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("parse date error. [YYYY.MM.DD]"))
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			return
		}
		next(w, r)
	}
}

func LogMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		end := time.Now()
		slog.Info(fmt.Sprintf("lead_time: %d | url: %s | method: %s", end.Sub(start).Milliseconds(), r.RemoteAddr, r.Method))
	}
}

type ErrorResponse struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func WriteError(messageError string, status int, w http.ResponseWriter) {
	errResponse := ErrorResponse{Err: messageError, Status: status}
	data, err := json.Marshal(errResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("server error"))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("server error"))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}
}
