package model

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User represents a row from 'public.users'.
type User struct {
	ID        uuid.UUID `json:"id"`         // id
	CreatedAt time.Time `json:"created_at"` // created_at
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	Name      string    `json:"name"`       // name
	Email     string    `json:"email"`      // email
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.users (` +
		`id, created_at, updated_at, name, email` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)`
	// run
	logf(sqlstr, u.ID, u.CreatedAt, u.UpdatedAt, u.Name, u.Email)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.CreatedAt, u.UpdatedAt, u.Name, u.Email); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.users SET ` +
		`created_at = $1, updated_at = $2, name = $3, email = $4 ` +
		`WHERE id = $5`
	// run
	logf(sqlstr, u.CreatedAt, u.UpdatedAt, u.Name, u.Email, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.CreatedAt, u.UpdatedAt, u.Name, u.Email, u.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.users (` +
		`id, created_at, updated_at, name, email` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at, name = EXCLUDED.name, email = EXCLUDED.email `
	// run
	logf(sqlstr, u.ID, u.CreatedAt, u.UpdatedAt, u.Name, u.Email)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.CreatedAt, u.UpdatedAt, u.Name, u.Email); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.users ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByID(ctx context.Context, db DB, id uuid.UUID) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, created_at, updated_at, name, email ` +
		`FROM public.users ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name, &u.Email); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}