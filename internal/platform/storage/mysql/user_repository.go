package mysql

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
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
	q, args, err := sq.Select(sqlUserColumns...).
		From(sqlUserTable).
		Where(sq.Eq{"id": userID.String()}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build the sql query: %v", err)
	}

	var dbUser sqlUser

	err = r.db.QueryRowContext(ctx, q, args...).Scan(
		&dbUser.ID,
		&dbUser.Name,
		&dbUser.Surnames,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.Country,
		&dbUser.Phone,
		&dbUser.PostalCode)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the user: %w", err)
	}

	user, err := user.NewUser(
		dbUser.ID,
		dbUser.Name,
		dbUser.Surnames,
		dbUser.Email,
		dbUser.Password,
		dbUser.Country,
		dbUser.Phone,
		dbUser.PostalCode)

	if err != nil {
		log.Errorf("Failed instanciating fetched user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, userEmail user.UserEmail) (*user.User, error) {
	q, args, err := sq.Select(sqlUserColumns...).
		From(sqlUserTable).
		Where(sq.Eq{"email": userEmail.String()}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build the sql query: %v", err)
	}

	var dbUser sqlUser

	err = r.db.QueryRowContext(ctx, q, args...).Scan(
		&dbUser.ID,
		&dbUser.Name,
		&dbUser.Surnames,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.Country,
		&dbUser.Phone,
		&dbUser.PostalCode)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the user: %w", err)
	}

	user, err := user.NewUser(
		dbUser.ID,
		dbUser.Name,
		dbUser.Surnames,
		dbUser.Email,
		dbUser.Password,
		dbUser.Country,
		dbUser.Phone,
		dbUser.PostalCode)

	if err != nil {
		log.Errorf("Failed instanciating fetched user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Save(ctx context.Context, user user.User) error {
	sqlUser := buildSQLUser(user)
	q, args, err := sq.Insert(sqlUserTable).
		Columns(sqlUserColumns...).
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
		log.Errorf("Error creating new user: %v", err)
		return fmt.Errorf("failed to insert new user: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user user.User) error {
	sqlUser := buildSQLUser(user)
	builder := sq.Update(sqlUserTable)
	q, args, err := buildUpdateArgument(sqlUser, builder).
		Where(sq.Eq{"id": user.ID().String()}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build the sql query: %v", err)
	}
	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update existing user: %w", err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userID user.UserID) error {
	q, args, err := sq.Delete(sqlUserTable).
		Where(sq.Eq{"id": userID.String()}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build the sql query: %v", err)
	}
	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to delete existing user: %w", err)
	}

	return nil
}

func buildSQLUser(user user.User) sqlUser {
	return sqlUser{
		ID:         user.ID().String(),
		Name:       user.Name().String(),
		Surnames:   user.Surnames().String(),
		Email:      user.Email().String(),
		Password:   user.Password().String(),
		Country:    user.Country().String(),
		Phone:      user.Phone().String(),
		PostalCode: user.PostalCode().String(),
	}
}

func buildUpdateArgument(sqlUser sqlUser, builder sq.UpdateBuilder) sq.UpdateBuilder {
	sqlUserMap := structs.Map(sqlUser)
	for _, column := range sqlUpdatableColumns {
		v, ok := sqlUserMap[column]
		if ok && v != "" {
			builder = builder.Set(column, v)
		}
	}
	return builder
}
