package ports

import (
	"nito/api/internals/core/domain"

	"github.com/gin-gonic/gin"
)

type ICharacterService interface {
	GetAllCharacters() []domain.Character
	CreateCharacter(character domain.Character) error
	FindById(id string) *domain.Character
}

type ICharacterRepository interface {
	GetAll() []domain.Character
	CreateCharacter(*domain.Character) error
	FindByName(Username *string) *domain.Character
	FindById(id *string) *domain.Character
}

type ICharacterHandler interface {
	GetAllCharacters(*gin.Context)
	CreateCharacter(c *gin.Context)
	FindById(c *gin.Context)
}
