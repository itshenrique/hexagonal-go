package ports

import (
	"nito/api/internals/core/domain"

	"github.com/gin-gonic/gin"
)

type IAccountService interface {
	Login(username string, password string) (*domain.Account, error)
	GetAllAccounts() []domain.Account
	CreateAccount(username string, password string, passwordConfirmation string) error
	FindById(id string) *domain.Account
}

type IAccountRepository interface {
	GetAll() []domain.Account
	CreateAccount(username *string, password *string) error
	FindByUsername(Username *string) *domain.Account
	FindById(id *string) *domain.Account
}

type IAccountHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetAllAccounts(*gin.Context)
	CreateAccount(c *gin.Context)
	FindById(c *gin.Context)
}
