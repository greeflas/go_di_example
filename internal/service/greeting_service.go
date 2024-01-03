package service

import (
	"fmt"
	"github.com/greeflas/go_di_example/internal/repository"
)

type GreetingService struct {
	userRepository    *repository.UserRepository
	messageRepository *repository.MessageRepository
}

func NewGreetingService(
	userRepository *repository.UserRepository,
	messageRepository *repository.MessageRepository,
) *GreetingService {
	return &GreetingService{
		userRepository:    userRepository,
		messageRepository: messageRepository,
	}
}

func (s *GreetingService) GetGreetingMessage(userId int) (string, error) {
	userName := "Guest"

	if userId > 0 {
		user, err := s.userRepository.GetById(userId)
		if err != nil {
			return "", err
		}

		userName = user.Name
	}

	message, err := s.messageRepository.GetMessage()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(message.Pattern, userName), nil
}
