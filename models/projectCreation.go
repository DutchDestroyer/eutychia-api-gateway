package models

type ProjectCreation struct {
	Project Project
	Participants []Participant
	Tests []string
}
