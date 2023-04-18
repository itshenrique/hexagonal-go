package ports

import "nito/api/internals/core/domain"

type ISessionService interface {
	Set(key string, value domain.Session) error
	Get(key string) (*domain.Session, error)
	Delete(key string) error
}

type ISessionRepository interface {
	Set(key string, value domain.Session, maxAge int) error
	Get(key string) (*domain.Session, error)
	Delete(key string) error
}
