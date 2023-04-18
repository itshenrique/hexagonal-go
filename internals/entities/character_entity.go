package entities

type Character struct {
	Base
	Name             string `gorm:"type:varchar(40);" json:"name"`
	Background       string `gorm:"type:varchar(255);" json:"background"`
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
