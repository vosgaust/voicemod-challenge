package user

import (
	"context"
	"fmt"

	"github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return UserService{userRepository}
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (user.User, error) {
	userEmail, err := user.NewUserEmail(email)
	if err != nil {
		return user.User{}, nil
	}
	foundUser, err := s.userRepository.FindByEmail(ctx, userEmail)
	if err != nil {
		return user.User{}, nil
	}

	return *foundUser, nil
}

func (s *UserService) Find(ctx context.Context, id string) (user.User, error) {
	userID, err := user.NewUserID(id)
	if err != nil {
		return user.User{}, nil
	}
	foundUser, err := s.userRepository.Find(ctx, userID)
	if err != nil {
		return user.User{}, nil
	}

	return *foundUser, nil
}

func (s *UserService) Create(ctx context.Context, id, name, surnames, email, password, country, phone, postalCode string) error {
	hashedPassword := user.EncryptPassword(password, id)
	user, err := user.NewUser(id, name, surnames, email, hashedPassword, country, phone, postalCode)
	if err != nil {
		fmt.Printf("Error initializing user: %\n", err)
		return err
	}

	err = s.userRepository.Save(ctx, user)
	if err != nil {
		fmt.Printf("Error saving user into repository: %v\n", err)
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
		fmt.Printf("Failed instanciating update user: %v\n", err)
		return err
	}

	existingUser, err := s.Find(ctx, id)
	fmt.Println(existingUser)
	if err != nil {
		fmt.Printf("Failed retrieving existing user to update: %v\n", err)
		return err
	}

	if existingUser.Password().String() != user.EncryptPassword(password, id) {
		fmt.Println("Old password is incorrect")
		fmt.Println(existingUser.Password().String())
		fmt.Println(user.EncryptPassword(password, id))
		return fmt.Errorf("incorrect current password")
	}

	err = s.userRepository.Update(ctx, updatedUser)
	if err != nil {
		fmt.Printf("Failed to persist user update: %v\n", err)
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
