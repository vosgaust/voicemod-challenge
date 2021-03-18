package user

import "context"

// UserRepository defines the expected behaviour from a user storage.
type UserRepository interface {
	Find(ctx context.Context, userID UserID) (*User, error)
	Save(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, userID UserID) error
}
