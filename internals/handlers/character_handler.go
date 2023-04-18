package handlers

import (
	"net/http"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/utils/http_error"

	"github.com/gin-gonic/gin"
)

// DI

type CharacterHandler struct {
	service        ports.ICharacterService
	sessionService ports.ISessionService
}

func NewCharacterHandler(service ports.ICharacterService) ports.ICharacterHandler {
	return &CharacterHandler{
		service: service,
	}
}

// Structs

type CreateCharacter struct {
	Name             string `json:"name"`
	Background       string `json:"background"`
	ExperiencePoints uint64 `json:"experiencePoints"`
	MaxHitPoints     uint32 `json:"maxHitPoints"`
	CurrentHitPoints uint32 `json:"currentHitPoints"`
	Strength         uint32 `json:"strength"`
	Dexterity        uint32 `json:"dexterity"`
	Constitution     uint32 `json:"constitution"`
	Intelligence     uint32 `json:"intelligence"`
	Wisdom           uint32 `json:"wisdom"`
	Charisma         uint32 `json:"charisma"`
}

// Functions

func (h *CharacterHandler) GetAllCharacters(c *gin.Context) {
	characters := h.service.GetAllCharacters()
	c.IndentedJSON(http.StatusOK, characters)
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	var body *CreateCharacter
	if err := c.ShouldBindJSON(&body); err != nil {
		error := http_error.NewBadRequestError("invalid json body")
		c.IndentedJSON(error.Status(), error)
		return
	}

	err := h.service.CreateCharacter(domain.Character{
		Name:             body.Name,
		Background:       body.Background,
		ExperiencePoints: body.ExperiencePoints,
		MaxHitPoints:     body.MaxHitPoints,
		CurrentHitPoints: body.CurrentHitPoints,
		Strength:         body.Strength,
		Dexterity:        body.Dexterity,
		Constitution:     body.Constitution,
		Intelligence:     body.Intelligence,
		Wisdom:           body.Wisdom,
		Charisma:         body.Charisma,
	})

	if err != nil {
		error := http_error.NewBadRequestError(err.Error())
		c.IndentedJSON(error.Status(), error)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h *CharacterHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	character := h.service.FindById(id)

	c.IndentedJSON(http.StatusOK, character)
}
