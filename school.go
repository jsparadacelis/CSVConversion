package main

type School struct {
	id         string
	name       string
	phone      string
	address    string
	city       string
	state      string
	zip        string
	schoolType string
	principal  string
	website    string
	other      string
	image      string
	active     string
}

func NewSchool(args ...string) *School {
	newSchool := School{
		id:         args[0],
		name:       args[1],
		phone:      args[2],
		address:    args[3],
		city:       args[4],
		state:      args[5],
		zip:        args[6],
		schoolType: args[7],
		principal:  args[8],
		website:    args[9],
		other:      args[10],
		image:      args[11],
		active:     args[12],
	}

	return &newSchool
}
