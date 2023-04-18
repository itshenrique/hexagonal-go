package domain

type Character struct {
	ID               string
	Name             string
	Background       string
	ExperiencePoints uint64
	MaxHitPoints     uint32
	CurrentHitPoints uint32
	Strength         uint32
	Dexterity        uint32
	Constitution     uint32
	Intelligence     uint32
	Wisdom           uint32
	Charisma         uint32
}
