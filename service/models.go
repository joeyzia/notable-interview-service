package service

// PhysicansRequest - The request for GET physicians
type PhysiciansRequest struct {

}

// PhysiciansResponse - The response from GET physicians
type PhysiciansResponse struct {
	Physicians []Person `json:"physicians"`
}

// ScheduleRequest - The request for GET phyician's schedule
type ScheduleRequest struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
}

// ScheduleResponse - The request for GET phyician's schedule
type ScheduleResponse struct {
	Schedule []Appointment `json:"appointments"`
}

// Appointment - Data Model for a specific appointment
type Appointment struct {
	ID int `json:"id"`
	Person Person `json:"name"`
	// Time time.Time `json:"time"`
	Time string `json:"time"`
	Kind string `json:"kind"`
}

// Person - representing a single person
type Person struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
}