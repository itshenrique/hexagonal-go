package services

import (
	"nito/api/internals/config"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
)

// DI

type SessionService struct {
	sessionRepository ports.ISessionRepository
}

func NewSessionService(sessionRepository ports.ISessionRepository) ports.ISessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

// Functions

func (r *SessionService) Set(key string, value domain.Session) error {
	config := config.LoadConfig(".")
	return r.sessionRepository.Set(key, value, config.SessionMaxAgeInDays*86400)
}

func (r *SessionService) Get(key string) (*domain.Session, error) {
	return r.sessionRepository.Get(key)
}

func (r *SessionService) Delete(key string) error {
	return r.sessionRepository.Delete(key)
}
