package model

type Attendee struct {
	ID         string
	Name       string
	Address    string
	Occupation string
	Reason     string
	index      int
}

type AttendeeList struct {
	Attendees []Attendee
}
