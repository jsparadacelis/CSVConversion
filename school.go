package main

type School struct {
	ID         string
	Name       string
	Phone      string
	Address    string
	City       string
	State      string
	Zip        string
	SchoolType string
	Principal  string
	Website    string
	Other      string
	Image      string
	Active     string
}

func NewSchool(args ...string) *School {
	newSchool := School{
		ID:         args[0],
		Name:       args[1],
		Phone:      args[2],
		Address:    args[3],
		City:       args[4],
		State:      args[5],
		Zip:        args[6],
		SchoolType: args[7],
		Principal:  args[8],
		Website:    args[9],
		Other:      args[10],
		Image:      args[11],
		Active:     args[12],
	}

	return &newSchool
}
