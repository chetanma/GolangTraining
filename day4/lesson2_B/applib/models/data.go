package models

import "errors"

type CustomerLevel int

const (
	Silver CustomerLevel = iota + 1
	Gold
	Platinum
)

type Customer struct {
	Name   string
	ID     int
	Points int
	Level  CustomerLevel
}

func NewCustomer(name string, id int, points int, level CustomerLevel) (*Customer, error) {

	if id < 1000 {
		return nil, errors.New("ID value too low")
	}
	return &Customer{
		Name:   name,
		ID:     id,
		Points: points,
		Level:  level,
	}, nil
}
