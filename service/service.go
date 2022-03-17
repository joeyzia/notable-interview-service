package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-kit/kit/log"
)

// Service - interface for API service
type Service interface {
	PhysiciansExecute(ctx context.Context, request PhysiciansRequest) (PhysiciansResponse, error)
	ScheduleExecute(ctx context.Context, request ScheduleRequest) (ScheduleResponse, error)
}

// service - the attributes related to the service
type service struct {
	mut      sync.Mutex
	logger   log.Logger
}

// NewService - Returns the new service's interface
func NewService(l log.Logger) Service {
	return &service{
		logger:   l,
	}
}

// PhysiciansExecute - Simple function to execute
func (s *service) PhysiciansExecute(ctx context.Context, request PhysiciansRequest) (resp PhysiciansResponse, err error) {
	log := log.With(s.logger, "event", "PhysiciansExecute")
	log.Log("PhysiciansExecute", "start")
	log.Log("mapPhysicians", "start")
	resp.Physicians = mapPhysicians();
	log.Log("mapPhysicians", "finished")

	log.Log("PhysiciansExecute", "finished")
	return resp, err
}

// ScheduleExecute - Simple function to execute
func (s *service) ScheduleExecute(ctx context.Context, request ScheduleRequest) (resp ScheduleResponse, err error) {
	log := log.With(s.logger, "event", "ScheduleExecute")
	log.Log("ScheduleExecute", "start")
	log.Log("mapPhysicianToSchedule", "start")
	data := mapPhysicianToSchedule()
	log.Log("mapPhysicianToSchedule", "finished")

	log.Log("data[key]", "finding")
	resp.Schedule = data[fmt.Sprintf("%s %s", request.FName, request.LName)]
	// There needs to be error handling if it doesn't exit -- with in memory, I wouldn't error handle with this model

	log.Log("data[key]", "found")

	log.Log("ScheduleExecute", "finished")
	return resp, err
}

// Helpers
func makePerson(fname, lname string) (person Person) {
	person.FName = fname;
	person.LName = lname;

	return person;
}

func makeSchedule(id int, person Person, time string, kind string) (app Appointment) {
	app.ID = id;
	app.Person = person;
	app.Time = time;
	app.Kind = kind;

	return app;
}

func mapPhysicians() []Person {
	var physicians []Person
	physicians = append(physicians, makePerson("Julius", "Hibbert"))
	physicians = append(physicians, makePerson("Algernop", "Krieger"))
	physicians = append(physicians, makePerson("Nick", "Riviera"))
	return physicians
}

func mapPhysicianToSchedule() map[string][]Appointment {
	schedule := make(map[string][]Appointment)
	schedule["Algernop Krieger"] = append(schedule["Algernop Krieger"], makeSchedule(1, makePerson("Sterling", "Archer"), "8:00AM", "New Patient"))
	schedule["Algernop Krieger"] = append(schedule["Algernop Krieger"], makeSchedule(2, makePerson("Cyril", "Figis"), "8:30AM", "Follow-up"))
		
	schedule["Julius Hibbert"] = append(schedule["Julius Hibbert"], makeSchedule(3, makePerson("Ray", "Gilette"), "9:00AM", "Follow-up"))

	schedule["Nick Riviera"] = append(schedule["Nick Rivera"], makeSchedule(4, makePerson("Ray", "Gilette"), "10:00AM", "New Patient"))
	return schedule
}