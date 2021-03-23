package user

import "context"

//go:generate mockery --case underscore --output ../../platform/storage/storagemocks --outpkg storagemocks --name UserRepository

// UserRepository defines the expected behaviour from a user storage.
type UserRepository interface {
	Find(ctx context.Context, userID UserID) (*User, error)
	FindByEmail(ctx context.Context, userEmail UserEmail) (*User, error)
	Save(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, userID UserID) error
}
