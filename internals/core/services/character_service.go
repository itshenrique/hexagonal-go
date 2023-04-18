package services

import (
	"errors"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
)

// DI

type CharacterService struct {
	characterRepository ports.ICharacterRepository
}

func NewCharacterService(characterRepository ports.ICharacterRepository) ports.ICharacterService {
	return &CharacterService{
		characterRepository: characterRepository,
	}
}

// Functions

func (s *CharacterService) GetAllCharacters() []domain.Character {
	characters := s.characterRepository.GetAll()

	return characters
}

func (s *CharacterService) CreateCharacter(character domain.Character) error {
	characterSearch := s.characterRepository.FindByName(&character.Name)

	if characterSearch.ID != "" {
		return errors.New("character name already registered")
	}

	return s.characterRepository.CreateCharacter(&character)
}

func (s *CharacterService) FindById(id string) *domain.Character {
	character := s.characterRepository.FindById(&id)

	if character.ID != "" {
		return character
	}

	return nil
}
