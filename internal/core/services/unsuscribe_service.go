package services

import "github.com/EdwinPirajan/Unsuscribe.git/internal/core/ports"

type UnsubscribeServiceImpl struct {
	repo ports.UnsubscribeRepository
}

func NewUnsubscribeService(repo ports.UnsubscribeRepository) *UnsubscribeServiceImpl {
	return &UnsubscribeServiceImpl{repo: repo}
}

func (s *UnsubscribeServiceImpl) Unsubscribe(email string) error {
	return s.repo.SaveUnsubscribe(email)
}
