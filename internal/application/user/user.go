package user

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

var ErrIncorrectPassword = errors.New("incorrect password")

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return UserService{userRepository}
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (user.User, error) {
	userEmail, err := user.NewUserEmail(email)
	if err != nil {
		return user.User{}, err
	}
	foundUser, err := s.userRepository.FindByEmail(ctx, userEmail)
	if err != nil {
		return user.User{}, err
	}

	return *foundUser, nil
}

func (s *UserService) Find(ctx context.Context, id string) (user.User, error) {
	userID, err := user.NewUserID(id)
	if err != nil {
		return user.User{}, err
	}
	foundUser, err := s.userRepository.Find(ctx, userID)
	if err != nil {
		return user.User{}, err
	}

	return *foundUser, nil
}

func (s *UserService) Create(ctx context.Context, id, name, surnames, email, password, country, phone, postalCode string) error {
	hashedPassword := user.EncryptPassword(password, id)
	user, err := user.NewUser(id, name, surnames, email, hashedPassword, country, phone, postalCode)
	if err != nil {
		log.Infof("Error initializing user: %v", err)
		return err
	}

	err = s.userRepository.Save(ctx, user)
	if err != nil {
		log.Infof("Error saving user into repository: %v", err)
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, id, name, surnames, email, password, newPassword, country, phone, postalCode string) error {
	newHashedPassword := user.EncryptPassword(newPassword, id)
	updatedUser, err := user.NewUser(
		id,
		name,
		surnames,
		email,
		newHashedPassword,
		country,
		phone,
		postalCode,
	)
	if err != nil {
		log.Infof("Failed instanciating update user: %v", err)
		return err
	}

	existingUser, err := s.Find(ctx, id)
	if err != nil {
		log.Infof("Failed retrieving existing user to update: %v", err)
		return err
	}

	if existingUser.Password().String() != user.EncryptPassword(password, id) {
		log.Info("Old password is incorrect")
		return ErrIncorrectPassword
	}

	err = s.userRepository.Update(ctx, updatedUser)
	if err != nil {
		log.Infof("Failed to persist user update: %v", err)
		return errors.New("failed to persist user update")
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	userID, err := user.NewUserID(id)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
