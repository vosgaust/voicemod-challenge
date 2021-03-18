package user

import (
	"context"

	"github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return UserService{userRepository}
}

func (s *UserService) Create(ctx context.Context, id, name, surnames, email, password, country, phone, postalCode string) error {
	user, err := user.NewUser(id, name, surnames, email, password, country, phone, postalCode)
	if err != nil {
		return err
	}

	err = s.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
