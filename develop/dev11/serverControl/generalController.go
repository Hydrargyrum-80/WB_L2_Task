package serverControl

import (
	"dev11/routers"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	router  http.Handler
	service *routers.EventControl
}

func NewController(service *routers.EventControl) *Controller {
	router := http.ServeMux{}
	controller := Controller{router: &router, service: service}
	router.HandleFunc("/ping", controller.Ping())
	router.HandleFunc("/event", LogMiddleware(controller.CheckDate(controller.EventHandler())))
	router.HandleFunc("/events_for_day", LogMiddleware(controller.CheckDate(controller.GetEventDay)))
	router.HandleFunc("/events_for_week", LogMiddleware(controller.CheckDate(controller.GetEventMonth)))
	router.HandleFunc("/events_for_month", LogMiddleware(controller.CheckDate(controller.GetEventWeek)))
	return &controller
}

func (c *Controller) Ping() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			WriteError("endpoint supports only GET method!", 404, w)
			return
		}
		w.WriteHeader(200)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}
}

func (c *Controller) EventHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.AddEvent(w, r)
		case http.MethodDelete:
			c.DeleteEvent(w, r)
		case http.MethodPut:
			c.PutEvent(w, r)
		default:
			WriteError("unsupport method!", 404, w)
			return
		}
	}
}

func (c *Controller) AddEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.AddEvent(r.Context(), t, r.URL.Query().Get("event"))
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.DeleteEvent(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) PutEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.UpdateEvent(r.Context(), t, r.URL.Query().Get("event"))
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) GetEventDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetDayEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func (c *Controller) GetEventWeek(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetWeekEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func (c *Controller) GetEventMonth(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetMonthEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func (c *Controller) GetRouter() http.Handler {
	return c.router
}
