package repositories

import (
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DI

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(dsn string) ports.IAccountRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect mariadb database")
	}

	db.Table("account").AutoMigrate(&entities.Account{})

	return &AccountRepository{
		db: db,
	}
}

// Functions

func (r *AccountRepository) GetAll() []domain.Account {
	var Accounts []domain.Account
	r.db.Table("account").Select("username", "id").Find(&Accounts)
	return Accounts
}

func (r *AccountRepository) CreateAccount(username *string, password *string) error {
	entity := entities.Account{
		Username: *username,
		Password: *password,
	}

	result := r.db.Table("account").Create(&entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AccountRepository) FindByUsername(Username *string) *domain.Account {
	var entity domain.Account
	r.db.Table("account").Where("username = ?", *Username).Find(&entity)
	return &entity
}

func (r *AccountRepository) FindById(id *string) *domain.Account {
	var entity entities.Account
	r.db.Table("account").First(&entity, "id = ?", *id)
	return &domain.Account{
		ID:       entity.ID,
		Username: entity.Username,
	}
}
