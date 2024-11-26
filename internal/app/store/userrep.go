package store

import "github.com/EWinterhalter/go_api/internal/app/model"

type Userrep struct {
	store *Store
}

func (r *Userrep) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, e_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.E_password,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}
func (r *Userrep) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, e_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.E_password,
	); err != nil {
		return nil, err
	}
	return u, nil
}
