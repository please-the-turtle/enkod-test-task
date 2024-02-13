package postgresql

import (
	"context"

	"github.com/gocraft/dbr/v2"
	"github.com/please-the-turtle/encod-test-task/app"
)

type postgresPersonRepository struct {
	session *dbr.Session
}

func NewPostgresPersonRepository(session *dbr.Session) app.PersonRepository {
	return &postgresPersonRepository{
		session: session,
	}
}

func (p *postgresPersonRepository) fetch(ctx context.Context, query string) (res []app.Person, err error) {
	rows, err := p.session.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	res = make([]app.Person, 0)
	for rows.Next() {
		t := app.Person{}
		err = rows.Scan(
			&t.Id,
			&t.FirstName,
			&t.LastName,
			&t.Email,
			&t.Phone,
		)

		if err != nil {
			return nil, err
		}

		res = append(res, t)
	}
	return
}

func (p *postgresPersonRepository) Fetch(ctx context.Context) (res []app.Person, err error) {
	query := `SELECT id,first_name,last_name,email,phone
						FROM person`

	res, err = p.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}
