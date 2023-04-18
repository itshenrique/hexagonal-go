package repositories

import (
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DI

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(dsn string) ports.ICharacterRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect mariadb database")
	}

	db.Table("character").AutoMigrate(&entities.Character{})

	return &CharacterRepository{
		db: db,
	}
}

// Functions

func (r *CharacterRepository) GetAll() []domain.Character {
	var Characters []domain.Character
	r.db.Table("character").Find(&Characters)
	return Characters
}

func (r *CharacterRepository) CreateCharacter(character *domain.Character) error {
	entity := entities.Character{
		Name:             character.Name,
		Background:       character.Background,
		ExperiencePoints: character.ExperiencePoints,
		MaxHitPoints:     character.MaxHitPoints,
		CurrentHitPoints: character.CurrentHitPoints,
		Strength:         character.Strength,
		Dexterity:        character.Dexterity,
		Constitution:     character.Constitution,
		Intelligence:     character.Intelligence,
		Wisdom:           character.Wisdom,
		Charisma:         character.Charisma,
	}

	result := r.db.Table("character").Create(&entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CharacterRepository) FindByName(Username *string) *domain.Character {
	var entity domain.Character
	r.db.Where("username = ?", Username).Find(&entity)
	return &entity
}

func (r *CharacterRepository) FindById(id *string) *domain.Character {
	var entity entities.Character
	r.db.First(&entity, "id = ?", id)
	return &domain.Character{
		ID:               entity.ID,
		Name:             entity.Name,
		Background:       entity.Background,
		ExperiencePoints: entity.ExperiencePoints,
		MaxHitPoints:     entity.MaxHitPoints,
		CurrentHitPoints: entity.CurrentHitPoints,
		Strength:         entity.Strength,
		Dexterity:        entity.Dexterity,
		Constitution:     entity.Constitution,
		Intelligence:     entity.Intelligence,
		Wisdom:           entity.Wisdom,
		Charisma:         entity.Charisma,
	}
}
