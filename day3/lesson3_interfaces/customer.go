package main

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

	if name == "" {
		return nil, errors.New("Empty name")
	}
	if id < 1000 {
		return nil, errors.New("ID value too low")
	}
	if points < 0 {
		return nil, errors.New("Points must be positive")
	}
	// if (level != 1 ) || (level != 2) || (level != 3) {
	// 	return nil, errors.New("CustomerLevel must be Silver, Gold or Platinum")
	// }
	return &Customer{
		Name:   name,
		ID:     id,
		Points: points,
		Level:  level,
	}, nil
}
