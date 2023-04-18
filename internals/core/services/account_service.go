package services

import (
	"errors"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/utils/crypto_util"
)

// DI

type AccountService struct {
	accountRepository ports.IAccountRepository
}

func NewAccountService(accountRepository ports.IAccountRepository) ports.IAccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

// Functions

func (s *AccountService) Login(username string, password string) (*domain.Account, error) {
	account := s.accountRepository.FindByUsername(&username)
	isValid := crypto_util.ComparePassword(&account.Password, &password)

	if !isValid {
		return nil, errors.New("user is not valid")
	}

	return account, nil
}

func (s *AccountService) GetAllAccounts() []domain.Account {
	accounts := s.accountRepository.GetAll()
	return accounts
}

func (s *AccountService) CreateAccount(username string, password string, passwordConfirmation string) error {
	account := s.accountRepository.FindByUsername(&username)
	if account.ID != "" {
		return errors.New("user already registered")
	}

	if password != passwordConfirmation {
		return errors.New("passwords are not equal")
	}

	acc := s.accountRepository.FindByUsername(&username)

	if acc.ID != "" {
		return nil
	}

	return s.accountRepository.CreateAccount(&username, &password)
}

func (s *AccountService) FindById(id string) *domain.Account {
	acc := s.accountRepository.FindById(&id)

	if acc.ID != "" {
		return &domain.Account{
			ID:       acc.ID,
			Username: acc.Username,
		}
	}
	return nil
}
