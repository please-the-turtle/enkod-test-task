package logic

import (
	"context"
	"time"

	"github.com/please-the-turtle/encod-test-task/app"
)

type PersonLogic struct {
	personRepo     app.PersonRepository
	contextTimeout time.Duration
}

func NewPersonLogic(repo app.PersonRepository, timeout time.Duration) app.PersonLogic {
	return &PersonLogic{
		personRepo:     repo,
		contextTimeout: timeout,
	}
}

func (p *PersonLogic) Fetch(ctx context.Context) (res []app.Person, err error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	res, err = p.personRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}
