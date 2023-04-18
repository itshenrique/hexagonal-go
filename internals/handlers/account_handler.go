package handlers

import (
	"net/http"
	"nito/api/internals/config"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/utils/crypto_util"
	"nito/api/internals/utils/http_error"

	"github.com/gin-gonic/gin"
)

// DI

type AccountHandler struct {
	service        ports.IAccountService
	sessionService ports.ISessionService
}

func NewAccountHandler(service ports.IAccountService, sessionService ports.ISessionService) ports.IAccountHandler {
	return &AccountHandler{
		service:        service,
		sessionService: sessionService,
	}
}

// Structs

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAccount struct {
	LoginUser
	ConfirmPassword string `json:"confirmPassword"`
}

// Functions

func (h *AccountHandler) Login(c *gin.Context) {
	config := config.LoadConfig(".")

	var body *LoginUser
	if err := c.ShouldBindJSON(&body); err != nil {
		error := http_error.NewBadRequestError("invalid json body")
		c.IndentedJSON(error.Status(), error)
		return
	}

	account, err := h.service.Login(body.Username, body.Password)

	if err != nil {
		error := http_error.NewBadRequestError(err.Error())
		c.IndentedJSON(error.Status(), error)
		return
	}

	encryptSessionId, err := crypto_util.Encrypt(account.ID, config.SessionSecret)

	if err != nil {
		error := http_error.NewInternalServerError("failed to create session", err)
		c.IndentedJSON(error.Status(), error)
		return
	}

	h.sessionService.Delete(account.ID)
	h.sessionService.Set(account.ID, domain.Session{
		ID:       account.ID,
		Username: account.Username,
	})

	c.SetCookie("session-id", *encryptSessionId, config.SessionMaxAgeInDays*86400, "", "", false, true)
}

func (h *AccountHandler) Logout(c *gin.Context) {
	config := config.LoadConfig(".")
	sessionID, err := c.Cookie("session-id")

	if err != nil {
		error := http_error.NewForbiddenError("invalid user information")
		c.JSON(error.Status(), error)
		return
	}

	decriptedSessionId, err := crypto_util.Decrypt(sessionID, config.SessionSecret)

	h.sessionService.Delete(*decriptedSessionId)

	c.SetCookie("session-id", "", -1, "", "", false, true)
}

func (h *AccountHandler) GetAllAccounts(c *gin.Context) {
	accounts := h.service.GetAllAccounts()
	c.IndentedJSON(http.StatusOK, accounts)
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var body *CreateAccount
	if err := c.ShouldBindJSON(&body); err != nil {
		error := http_error.NewBadRequestError("invalid json body")
		c.IndentedJSON(error.Status(), error)
		return
	}

	err := h.service.CreateAccount(body.Username, body.Password, body.ConfirmPassword)

	if err != nil {
		error := http_error.NewBadRequestError(err.Error())
		c.IndentedJSON(error.Status(), error)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h *AccountHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	account := h.service.FindById(id)

	c.IndentedJSON(http.StatusOK, account)
}
