package app

import "context"

type Person struct {
	Id        int64
	Email     string
	Phone     string
	FirstName string
	LastName  string
}

type PersonLogic interface {
	Fetch(ctx context.Context) (res []Person, err error)
}

type PersonRepository interface {
	Fetch(ctx context.Context) (res []Person, err error)
}
