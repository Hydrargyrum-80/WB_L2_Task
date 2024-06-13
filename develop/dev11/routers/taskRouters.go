package routers

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type EventControl struct {
	sync.Map
}

var (
	ErrEventAlreadyExists = errors.New("error: event already exists")
	ErrEventNotExists     = errors.New("error: event not exists")
)

func NewEventControl() *EventControl {
	return &EventControl{sync.Map{}}
}

func (s *EventControl) AddEvent(ctx context.Context, date time.Time, event string) error {
	dateStr := date.Format(time.DateOnly)
	_, ok := s.Load(dateStr)
	if ok {
		return ErrEventAlreadyExists
	}
	s.Store(dateStr, event)
	fmt.Println(s.Load(dateStr))
	return nil
}

func (s *EventControl) UpdateEvent(ctx context.Context, date time.Time, event string) error {
	dateString := date.Format(time.DateOnly)
	_, ok := s.Load(dateString)
	if !ok {
		return ErrEventNotExists
	}
	s.Store(dateString, event)
	return nil
}

func (s *EventControl) DeleteEvent(ctx context.Context, date time.Time) error {
	dateString := date.Format(time.DateOnly)
	_, ok := s.Load(dateString)
	if !ok {
		return ErrEventNotExists
	}
	s.Delete(dateString)
	return nil
}

type EventResponse struct {
	Date  string `json:"date"`
	Event string `json:"event"`
}

func (s *EventControl) GetDayEvents(ctx context.Context, date time.Time) (*EventResponse, error) {
	dateString := date.Format(time.DateOnly)
	value, ok := s.Load(dateString)
	if !ok {
		return nil, ErrEventNotExists
	}
	return &EventResponse{Date: dateString, Event: value.(string)}, nil
}

func (s *EventControl) GetWeekEvents(ctx context.Context, date time.Time) ([]EventResponse, error) {
	var events []EventResponse
	for i := 0; i < 7; i++ {
		dateString := date.Format(time.DateOnly)
		value, ok := s.Load(dateString)
		if ok {
			events = append(events, EventResponse{Date: dateString, Event: value.(string)})
		}
		date = date.Add(time.Hour * 24)
	}
	if len(events) == 0 {
		return nil, ErrEventNotExists
	}
	return events, nil
}

func (s *EventControl) GetMonthEvents(ctx context.Context, date time.Time) ([]EventResponse, error) {
	var events []EventResponse
	for i := 0; i < 31; i++ {
		dateString := date.Format(time.DateOnly)
		value, ok := s.Load(dateString)
		if ok {
			events = append(events, EventResponse{Date: dateString, Event: value.(string)})
		}
		date = date.Add(time.Hour * 24)
	}
	if len(events) == 0 {
		return nil, ErrEventNotExists
	}
	return events, nil
}
