package mysql

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

// UserRepository is a MySQL implementation for UserRepository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a MySQL-based implementation of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Find(ctx context.Context, userID user.UserID) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) Save(ctx context.Context, user user.User) error {
	sqlUser := sqlUser{
		ID:         user.ID().String(),
		Name:       user.Name().String(),
		Surnames:   user.Surnames().String(),
		Email:      user.Email().String(),
		Password:   user.Password().String(),
		Country:    user.Country().String(),
		Phone:      user.Phone().String(),
		PostalCode: user.PostalCode().String(),
	}

	q, args, err := sq.Insert(sqlUserTable).
		Columns(sqlInsertUserColumns...).
		Values(
			sqlUser.ID,
			sqlUser.Name,
			sqlUser.Surnames,
			sqlUser.Email,
			sqlUser.Password,
			sqlUser.Country,
			sqlUser.Phone,
			sqlUser.PostalCode,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build the sql query: %v", err)
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user user.User) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userID user.UserID) error {
	return nil
}
