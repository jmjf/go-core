package models

type Animal struct {
	Id         int
	Family     string
	Genus      string
	Species    string
	CommonName string
}

var (
	animal []*Animal
	nextId = 1
)
