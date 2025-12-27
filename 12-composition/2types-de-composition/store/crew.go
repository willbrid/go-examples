package store

type Crew struct {
	Captain, FirstOfficer string
}

func NewCrew(captain, firstOffice string) *Crew {
	return &Crew{captain, firstOffice}
}
